package randutil

import (
	"io"
	"sync"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

// Root derives domain-separated entropy streams for a workspace.
type Root interface {
	Derive(label string) (core.Source, error)
}

// SecureRoot returns a Root seeded from crypto/rand.
func SecureRoot() Root {
	return SecureRootWithSource(nil)
}

// SecureRootWithSource returns a Root seeded from src. If src is nil,
// crypto/rand.Reader is used.
func SecureRootWithSource(src core.Source) Root {
	if src == nil {
		src = adapters.CryptoSource()
	}
	return &secureRoot{src: src}
}

// DeterministicRoot returns a Root derived from seed.
// WARNING: Deterministic roots are intended for tests and replayable
// simulations. DO NOT USE FOR TOKENS / AUTH unless the seed is high-entropy
// and kept secret.
func DeterministicRoot(seed []byte) Root {
	return newDeterministicRoot(seed)
}

type secureRoot struct {
	src core.Source

	once sync.Once
	seed [32]byte
	err  error

	mu     sync.RWMutex
	closed bool
}

func (r *secureRoot) Derive(label string) (core.Source, error) {
	r.once.Do(func() {
		var seed [32]byte
		_, err := io.ReadFull(r.src, seed[:])
		r.mu.Lock()
		defer r.mu.Unlock()
		r.err = err
		if err == nil {
			r.seed = seed
		}
	})

	r.mu.RLock()
	closed := r.closed
	err := r.err
	seed := r.seed
	r.mu.RUnlock()

	if closed {
		return nil, core.ErrSourceClosed
	}
	if err != nil {
		return nil, err
	}
	return adapters.DeriveSource(seed[:], label)
}

func (r *secureRoot) Close() error {
	r.mu.Lock()
	if r.closed {
		r.mu.Unlock()
		return nil
	}
	r.closed = true
	core.Zero(r.seed[:])
	src := r.src
	r.mu.Unlock()
	if closer, ok := src.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

type seedRoot struct {
	mu     sync.RWMutex
	seed   []byte
	closed bool
}

func newSeedRoot(seed []byte) *seedRoot {
	copied := make([]byte, len(seed))
	copy(copied, seed)
	return &seedRoot{seed: copied}
}

func (r *seedRoot) Derive(label string) (core.Source, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	if r.closed {
		return nil, core.ErrSourceClosed
	}
	return adapters.DeriveSource(r.seed, label)
}

func (r *seedRoot) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	r.closed = true
	core.Zero(r.seed)
	return nil
}
