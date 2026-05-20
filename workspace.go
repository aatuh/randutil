package randutil

import (
	"errors"
	"io"
	"sync"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/dist"
	"github.com/aatuh/randutil/v2/email"
	"github.com/aatuh/randutil/v2/nanoid"
	"github.com/aatuh/randutil/v2/numeric"
	"github.com/aatuh/randutil/v2/randstring"
	"github.com/aatuh/randutil/v2/randtime"
	"github.com/aatuh/randutil/v2/ulid"
	"github.com/aatuh/randutil/v2/uuid"
)

const (
	defaultMaxCachedStreams = 64
	subRootLabelPrefix      = "randutil workspace subroot v1 "
)

// UsageHook receives byte deltas for a label.
type UsageHook func(label string, delta uint64)

type streamEntry struct {
	gen     *core.Generator
	counter *adapters.CountingSource
}

// Workspace routes random streams by name from a shared root seed.
// Stream outputs are deterministic for a given root+label.
// Sub derives nested workspaces, Usage tracks bytes per cached label,
// and Close attempts to zero internal state.
//
// Concurrency: Stream and Rand are safe for concurrent use.
type Workspace struct {
	root      Root
	maxCached int
	usageHook UsageHook

	mu      sync.Mutex
	streams map[string]*streamEntry
	order   []string
	closed  bool
}

// WorkspaceOptions configure stream caching.
type WorkspaceOptions struct {
	// MaxCached is the maximum number of cached streams. Use a negative value
	// to disable caching. If zero, a default is applied.
	MaxCached int

	// UsageHook receives byte deltas per label when streams are read.
	UsageHook UsageHook
}

// NewWorkspace returns a workspace using the provided root.
// If root is nil, SecureRoot is used.
func NewWorkspace(root Root) *Workspace {
	return NewWorkspaceWithOptions(root, WorkspaceOptions{})
}

// NewWorkspaceWithOptions returns a workspace with cache limits.
func NewWorkspaceWithOptions(root Root, opts WorkspaceOptions) *Workspace {
	if root == nil {
		root = SecureRoot()
	}
	maxCached := opts.MaxCached
	if maxCached == 0 {
		maxCached = defaultMaxCachedStreams
	}
	if maxCached < 0 {
		maxCached = 0
	}
	streams := map[string]*streamEntry{}
	if maxCached == 0 {
		streams = nil
	}
	return &Workspace{
		root:      root,
		maxCached: maxCached,
		usageHook: opts.UsageHook,
		streams:   streams,
	}
}

// Stream returns a cached RNG stream for label.
func (w *Workspace) Stream(label string) (*core.Generator, error) {
	if w == nil {
		return nil, errors.New("randutil: workspace is nil")
	}
	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		return nil, core.ErrWorkspaceClosed
	}
	if w.maxCached != 0 {
		if entry, ok := w.streams[label]; ok {
			gen := entry.gen
			w.mu.Unlock()
			return gen, nil
		}
	}
	w.mu.Unlock()

	src, err := w.root.Derive(label)
	if err != nil {
		return nil, err
	}
	var hook func(delta uint64)
	if w.usageHook != nil {
		labelCopy := label
		hook = func(delta uint64) {
			w.usageHook(labelCopy, delta)
		}
	}
	counter := adapters.NewCountingSource(src, hook)
	gen := core.New(counter)

	if w.maxCached == 0 {
		return gen, nil
	}

	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		_ = gen.Close()
		return nil, core.ErrWorkspaceClosed
	}
	if existing, ok := w.streams[label]; ok {
		w.mu.Unlock()
		_ = gen.Close()
		return existing.gen, nil
	}
	if w.maxCached > 0 && len(w.streams) >= w.maxCached {
		w.evictOldestLocked()
	}
	w.streams[label] = &streamEntry{gen: gen, counter: counter}
	w.order = append(w.order, label)
	w.mu.Unlock()
	return gen, nil
}

// Rand returns a convenience bundle of generators bound to label.
func (w *Workspace) Rand(label string) (Rand, error) {
	gen, err := w.Stream(label)
	if err != nil {
		return Rand{}, err
	}
	return Rand{
		Core:    gen,
		Numeric: numeric.New(gen),
		Dist:    dist.New(gen),
		String:  randstring.New(gen),
		UUID:    uuid.New(gen),
		Time:    randtime.New(gen),
		Email:   email.New(gen),
		NanoID:  nanoid.New(gen),
		ULID:    ulid.New(gen),
	}, nil
}

// Sub returns a nested workspace derived from label.
func (w *Workspace) Sub(label string) (*Workspace, error) {
	if w == nil {
		return nil, errors.New("randutil: workspace is nil")
	}
	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		return nil, core.ErrWorkspaceClosed
	}
	opts := WorkspaceOptions{
		MaxCached: w.maxCached,
		UsageHook: w.usageHook,
	}
	w.mu.Unlock()

	src, err := w.root.Derive(subRootLabelPrefix + label)
	if err != nil {
		return nil, err
	}
	var seed [32]byte
	if _, err := io.ReadFull(src, seed[:]); err != nil {
		if closer, ok := src.(io.Closer); ok {
			_ = closer.Close()
		}
		core.Zero(seed[:])
		return nil, err
	}
	if closer, ok := src.(io.Closer); ok {
		_ = closer.Close()
	}
	root := newSeedRoot(seed[:])
	core.Zero(seed[:])
	return NewWorkspaceWithOptions(root, opts), nil
}

// Usage returns the number of bytes read for a cached label.
func (w *Workspace) Usage(label string) (uint64, bool) {
	if w == nil {
		return 0, false
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.streams == nil || w.closed {
		return 0, false
	}
	entry, ok := w.streams[label]
	if !ok || entry == nil || entry.counter == nil {
		return 0, false
	}
	return entry.counter.Count(), true
}

// UsageSnapshot returns a copy of usage counters for cached labels.
func (w *Workspace) UsageSnapshot() map[string]uint64 {
	if w == nil {
		return map[string]uint64{}
	}
	w.mu.Lock()
	defer w.mu.Unlock()
	if w.streams == nil {
		return map[string]uint64{}
	}
	out := make(map[string]uint64, len(w.streams))
	for label, entry := range w.streams {
		if entry == nil || entry.counter == nil {
			continue
		}
		out[label] = entry.counter.Count()
	}
	return out
}

// Close zeroes workspace state and closes derived streams where possible.
func (w *Workspace) Close() error {
	if w == nil {
		return errors.New("randutil: workspace is nil")
	}
	w.mu.Lock()
	if w.closed {
		w.mu.Unlock()
		return nil
	}
	w.closed = true
	streams := w.streams
	w.streams = nil
	w.order = nil
	w.mu.Unlock()

	var firstErr error
	for _, entry := range streams {
		if entry == nil || entry.gen == nil {
			continue
		}
		if err := entry.gen.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	if closer, ok := w.root.(interface{ Close() error }); ok {
		if err := closer.Close(); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (w *Workspace) evictOldestLocked() {
	if len(w.order) == 0 {
		return
	}
	oldest := w.order[0]
	entry := w.streams[oldest]
	delete(w.streams, oldest)
	if entry != nil && entry.gen != nil {
		_ = entry.gen.Close()
	}
	copy(w.order, w.order[1:])
	w.order = w.order[:len(w.order)-1]
}
