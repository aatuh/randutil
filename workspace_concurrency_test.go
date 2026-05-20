package randutil

import (
	"errors"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/aatuh/randutil/v2/core"
)

type serialRoot struct {
	active atomic.Int32
	calls  atomic.Int32
}

func (r *serialRoot) Derive(string) (core.Source, error) {
	if got := r.active.Add(1); got != 1 {
		r.active.Add(-1)
		return nil, errConcurrentAccess{}
	}
	time.Sleep(time.Millisecond)
	r.calls.Add(1)
	r.active.Add(-1)
	return &serialSource{}, nil
}

type serialSource struct {
	active atomic.Int32
}

func (s *serialSource) Read(p []byte) (int, error) {
	if got := s.active.Add(1); got != 1 {
		s.active.Add(-1)
		return 0, errConcurrentAccess{}
	}
	time.Sleep(time.Millisecond)
	for i := range p {
		p[i] = byte(i)
	}
	s.active.Add(-1)
	return len(p), nil
}

type errConcurrentAccess struct{}

func (errConcurrentAccess) Error() string {
	return "concurrent access"
}

type nilSourceRoot struct{}

func (nilSourceRoot) Derive(string) (core.Source, error) {
	//nolint:nilnil // This test double covers a custom root that violates the contract.
	return nil, nil
}

type blockingRoot struct {
	started chan struct{}
	release chan struct{}
	once    sync.Once
}

func newBlockingRoot() *blockingRoot {
	return &blockingRoot{
		started: make(chan struct{}),
		release: make(chan struct{}),
	}
}

func (r *blockingRoot) Derive(string) (core.Source, error) {
	r.once.Do(func() {
		close(r.started)
	})
	<-r.release
	return &serialSource{}, nil
}

func (r *blockingRoot) Close() error {
	return nil
}

func TestWorkspaceSerializesRootDerive(t *testing.T) {
	root := &serialRoot{}
	ws := NewWorkspace(root)

	const workers = 16
	var wg sync.WaitGroup
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func(label string) {
			defer wg.Done()
			_, err := ws.Stream(label)
			errs <- err
		}(string(rune('a' + i)))
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Fatalf("Stream error: %v", err)
		}
	}
	if got := root.calls.Load(); got != workers {
		t.Fatalf("Derive calls = %d, want %d", got, workers)
	}
}

func TestWorkspaceRejectsNilDerivedSource(t *testing.T) {
	ws := NewWorkspace(nilSourceRoot{})
	if _, err := ws.Stream("alpha"); !errors.Is(err, core.ErrSourceClosed) {
		t.Fatalf("Stream error = %v, want ErrSourceClosed", err)
	}
	if _, err := ws.Sub("alpha"); !errors.Is(err, core.ErrSourceClosed) {
		t.Fatalf("Sub error = %v, want ErrSourceClosed", err)
	}
}

func TestWorkspaceDisabledCacheStreamDoesNotEscapeAfterClose(t *testing.T) {
	root := newBlockingRoot()
	ws := NewWorkspaceWithOptions(root, WorkspaceOptions{MaxCached: -1})

	streamErr := make(chan error, 1)
	go func() {
		gen, err := ws.Stream("alpha")
		if gen != nil {
			_ = gen.Close()
		}
		streamErr <- err
	}()
	<-root.started

	closeErr := make(chan error, 1)
	go func() {
		closeErr <- ws.Close()
	}()
	waitWorkspaceClosed(t, ws)
	close(root.release)

	if err := <-streamErr; !errors.Is(err, core.ErrWorkspaceClosed) {
		t.Fatalf("Stream error = %v, want ErrWorkspaceClosed", err)
	}
	if err := <-closeErr; err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestWorkspaceSubDoesNotEscapeAfterClose(t *testing.T) {
	root := newBlockingRoot()
	ws := NewWorkspace(root)

	subErr := make(chan error, 1)
	go func() {
		sub, err := ws.Sub("alpha")
		if sub != nil {
			_ = sub.Close()
		}
		subErr <- err
	}()
	<-root.started

	closeErr := make(chan error, 1)
	go func() {
		closeErr <- ws.Close()
	}()
	waitWorkspaceClosed(t, ws)
	close(root.release)

	if err := <-subErr; !errors.Is(err, core.ErrWorkspaceClosed) {
		t.Fatalf("Sub error = %v, want ErrWorkspaceClosed", err)
	}
	if err := <-closeErr; err != nil {
		t.Fatalf("Close error: %v", err)
	}
}

func TestWorkspaceSerializesStreamSourceReads(t *testing.T) {
	ws := NewWorkspace(&serialRoot{})
	gen, err := ws.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}

	const workers = 16
	var wg sync.WaitGroup
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := gen.Bytes(8)
			errs <- err
		}()
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Fatalf("Bytes error: %v", err)
		}
	}
}

func TestWorkspaceSerializesUsageHook(t *testing.T) {
	var active atomic.Int32
	var concurrent atomic.Bool
	ws := NewWorkspaceWithOptions(&serialRoot{}, WorkspaceOptions{
		UsageHook: func(string, uint64) {
			if got := active.Add(1); got != 1 {
				concurrent.Store(true)
			}
			time.Sleep(time.Millisecond)
			active.Add(-1)
		},
	})
	gen, err := ws.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}

	const workers = 16
	var wg sync.WaitGroup
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := gen.Bytes(8)
			errs <- err
		}()
	}
	wg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Fatalf("Bytes error: %v", err)
		}
	}
	if concurrent.Load() {
		t.Fatalf("UsageHook ran concurrently")
	}
}

func waitWorkspaceClosed(t *testing.T, ws *Workspace) {
	t.Helper()

	deadline := time.After(2 * time.Second)
	tick := time.NewTicker(time.Millisecond)
	defer tick.Stop()
	for {
		ws.mu.Lock()
		closed := ws.closed
		ws.mu.Unlock()
		if closed {
			return
		}
		select {
		case <-deadline:
			t.Fatalf("workspace was not closed before deadline")
		case <-tick.C:
		}
	}
}
