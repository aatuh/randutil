//go:build randutil_must
// +build randutil_must

package deterministic

import "github.com/aatuh/randutil/v2/core"

// MustSource returns a deterministic source or panics.
func MustSource(seed []byte) core.Source {
	src, err := Source(seed)
	if err != nil {
		panic(err)
	}
	return src
}

// MustSourceWithLabel returns a deterministic source or panics.
func MustSourceWithLabel(seed []byte, label string) core.Source {
	src, err := SourceWithLabel(seed, label)
	if err != nil {
		panic(err)
	}
	return src
}

// MustRNG returns a deterministic RNG or panics.
func MustRNG(seed []byte) core.RNG {
	rng, err := RNG(seed)
	if err != nil {
		panic(err)
	}
	return rng
}

// MustRNGWithLabel returns a deterministic RNG or panics.
func MustRNGWithLabel(seed []byte, label string) core.RNG {
	rng, err := RNGWithLabel(seed, label)
	if err != nil {
		panic(err)
	}
	return rng
}
