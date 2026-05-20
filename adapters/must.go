//go:build randutil_must
// +build randutil_must

package adapters

import "github.com/aatuh/randutil/v2/core"

// MustDeterministicSource returns a deterministic source or panics.
func MustDeterministicSource(seed []byte) core.Source {
	src, err := DeterministicSource(seed)
	if err != nil {
		panic(err)
	}
	return src
}

// MustDeterministicSourceWithLabel returns a deterministic source or panics.
func MustDeterministicSourceWithLabel(seed []byte, label string) core.Source {
	src, err := DeterministicSourceWithLabel(seed, label)
	if err != nil {
		panic(err)
	}
	return src
}

// MustDeriveSource returns a derived source or panics.
func MustDeriveSource(seed []byte, label string) core.Source {
	src, err := DeriveSource(seed, label)
	if err != nil {
		panic(err)
	}
	return src
}

// MustDeriveRNG returns a derived RNG or panics.
func MustDeriveRNG(seed []byte, label string) core.RNG {
	rng, err := DeriveRNG(seed, label)
	if err != nil {
		panic(err)
	}
	return rng
}

// MustFastSource returns a fast derived source or panics.
func MustFastSource() core.Source {
	src, err := FastSource()
	if err != nil {
		panic(err)
	}
	return src
}

// MustFastSourceWithSource returns a fast derived source or panics.
func MustFastSourceWithSource(src core.Source) core.Source {
	derived, err := FastSourceWithSource(src)
	if err != nil {
		panic(err)
	}
	return derived
}

// MustFastRNG returns a fast derived RNG or panics.
func MustFastRNG() core.RNG {
	rng, err := FastRNG()
	if err != nil {
		panic(err)
	}
	return rng
}
