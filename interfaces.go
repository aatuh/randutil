package randutil

import "github.com/aatuh/randutil/v2/core"

// Streamer provides access to a named generator stream.
type Streamer interface {
	Stream(label string) (*core.Generator, error)
}

// RandProvider provides access to a named Rand bundle.
type RandProvider interface {
	Rand(label string) (Rand, error)
}
