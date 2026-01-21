package randutil

import (
	"github.com/aatuh/randutil/v2/collection"
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/dist"
	"github.com/aatuh/randutil/v2/email"
	"github.com/aatuh/randutil/v2/numeric"
	"github.com/aatuh/randutil/v2/randstring"
	"github.com/aatuh/randutil/v2/randtime"
	"github.com/aatuh/randutil/v2/uuid"
)

// Rand provides access to generators from the subpackages, bound to a single
// entropy source. This eliminates the need to duplicate all the methods from
// each subpackage.
type Rand struct {
	// Core provides basic random number generation primitives.
	Core *core.Generator

	// Numeric provides numeric helpers (ranges, ints, bytes).
	Numeric *numeric.Generator

	// Dist provides statistical distributions.
	Dist *dist.Generator

	// String provides random string and token generation.
	String *randstring.Generator

	// UUID provides UUID v4 and v7 generation.
	UUID *uuid.Generator

	// Time provides random datetime generation.
	Time *randtime.Generator

	// Email provides random email address generation.
	Email *email.Generator
}

// New returns a Rand with all generators bound to src. Pass nil to use
// crypto/rand.
//
// Parameters:
//   - src: The entropy source to use.
//
// Returns:
//   - Rand: A new Rand with all generators bound to src.
func New(src core.Source) Rand {
	coreGen := core.New(src)
	return Rand{
		Core:    coreGen,
		Numeric: numeric.New(coreGen),
		Dist:    dist.New(coreGen),
		String:  randstring.New(coreGen),
		UUID:    uuid.New(coreGen),
		Time:    randtime.New(coreGen),
		Email:   email.New(coreGen),
	}
}

// Default returns a Rand using crypto/rand.
//
// Returns:
//   - Rand: A new Rand using crypto/rand.
func Default() Rand { return New(nil) }

// Secure is an alias for Default, provided for clarity when the
// caller wants to emphasize security properties.
//
// Returns:
//   - Rand: A new Rand using crypto/rand.
func Secure() Rand { return Default() }

// Source exposes the underlying entropy source.
//
// Returns:
//   - core.Source: The underlying entropy source.
func (r Rand) Source() core.Source { return r.Core.Source() }

// Collection returns a collection generator bound to this Rand's RNG.
func Collection[T any](r Rand) *collection.Generator[T] {
	if r.Core == nil {
		return collection.New[T](nil)
	}
	return collection.New[T](r.Core)
}
