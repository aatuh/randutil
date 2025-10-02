package randutil

import (
	"io"

	"github.com/aatuh/randutil/v2/collection"
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/email"
	"github.com/aatuh/randutil/v2/randstring"
	"github.com/aatuh/randutil/v2/randtime"
	"github.com/aatuh/randutil/v2/uuid"
)

// Rand provides access to all default generators from the subpackages,
// bound to a single entropy source. This eliminates the need to
// duplicate all the methods from each subpackage.
type Rand struct {
	// Core provides basic random number generation primitives.
	Core *core.Generator

	// String provides random string and token generation.
	String *randstring.Generator

	// UUID provides UUID v4 and v7 generation.
	UUID *uuid.Generator

	// Collection provides slice operations like shuffle and sampling.
	Collection *collection.Generator

	// Time provides random datetime generation.
	Time *randtime.Generator

	// Email provides random email address generation.
	Email *email.Generator
}

// New returns a Rand with all generators bound to src. Pass nil to use
// crypto/rand via the package default.
//
// Parameters:
//   - src: The entropy source to use.
//
// Returns:
//   - Rand: A new Rand with all generators bound to src.
func New(src io.Reader) Rand {
	coreGen := core.New(src)
	return Rand{
		Core:       coreGen,
		String:     randstring.New(src),
		UUID:       uuid.New(src),
		Collection: collection.New(src),
		Time:       randtime.New(src),
		Email:      email.New(src),
	}
}

// Default returns a Rand using the current secure source
// (crypto/rand by default).
//
// Returns:
//   - Rand: A new Rand using the current secure source.
func Default() Rand { return New(nil) }

// Secure is an alias for Default, provided for clarity when the
// caller wants to emphasize security properties.
//
// Returns:
//   - Rand: A new Rand using the current secure source.
func Secure() Rand { return Default() }

// Reader exposes the underlying entropy source.
//
// Returns:
//   - io.Reader: The current entropy source.
func (r Rand) Reader() io.Reader { return r.Core }

// Global convenience accessors.

// String provides access to the default string generator.
var String = randstring.Default

// UUID provides access to the default UUID generator.
var UUID = uuid.Default

// Collection provides access to the default collection generator.
var Collection = collection.Default

// Time provides access to the default time generator.
var Time = randtime.Default

// Email provides access to the default email generator.
var Email = email.Default

// Core provides access to the default core generator.
var Core = core.New(nil)
