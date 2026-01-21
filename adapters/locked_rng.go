package adapters

import (
	"sync"

	"github.com/aatuh/randutil/v2/core"
)

type lockedRNG struct {
	mu  sync.Mutex
	rng core.RNG
}

// LockedRNG returns an RNG wrapper that serializes access to rng.
// If rng is nil, it returns nil.
func LockedRNG(rng core.RNG) core.RNG {
	if rng == nil {
		return nil
	}
	return &lockedRNG{rng: rng}
}

func (l *lockedRNG) Read(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Read(p)
}

func (l *lockedRNG) Fill(p []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Fill(p)
}

func (l *lockedRNG) Bytes(n int) ([]byte, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Bytes(n)
}

func (l *lockedRNG) Uint64() (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Uint64()
}

func (l *lockedRNG) Uint64n(n uint64) (uint64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Uint64n(n)
}

func (l *lockedRNG) Intn(n int) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Intn(n)
}

func (l *lockedRNG) Int64n(n int64) (int64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Int64n(n)
}

func (l *lockedRNG) IntRange(minInclusive, maxInclusive int) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.IntRange(minInclusive, maxInclusive)
}

func (l *lockedRNG) Int32Range(minInclusive, maxInclusive int32) (int32, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Int32Range(minInclusive, maxInclusive)
}

func (l *lockedRNG) Int64Range(minInclusive, maxInclusive int64) (int64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Int64Range(minInclusive, maxInclusive)
}

func (l *lockedRNG) Float64() (float64, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Float64()
}

func (l *lockedRNG) Bool() (bool, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.rng.Bool()
}
