package numeric

import "github.com/aatuh/randutil/v2/core"

// Generator provides numeric random helpers backed by a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng core.RNG
}

// New returns a numeric Generator. If rng is nil, crypto/rand is used.
func New(rng core.RNG) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator{rng: rng}
}

// NewWithSource returns a numeric Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return New(core.New(src))
}
