package nanoid

import (
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/randstring"
)

// Generator builds NanoID strings using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	strings *randstring.Generator
}

// New returns a NanoID Generator. If rng is nil, crypto/rand is used.
func New(rng rng) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator{strings: randstring.New(rng)}
}

// NewWithSource returns a NanoID Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return New(core.New(src))
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// ID returns a NanoID string of the requested length.
func (g *Generator) ID(length int) (string, error) {
	return g.strings.StringWithCharset(length, DefaultAlphabet)
}
