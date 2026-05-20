//go:build randutil_must
// +build randutil_must

package randutil

import "github.com/aatuh/randutil/v2/core"

// MustDerive returns a derived Rand or panics.
func MustDerive(seed []byte, label string) Rand {
	r, err := Derive(seed, label)
	if err != nil {
		panic(err)
	}
	return r
}

// MustDeriveRNG returns a derived RNG or panics.
func MustDeriveRNG(seed []byte, label string) core.RNG {
	rng, err := DeriveRNG(seed, label)
	if err != nil {
		panic(err)
	}
	return rng
}

// MustFast returns a fast derived Rand or panics.
func MustFast() Rand {
	r, err := Fast()
	if err != nil {
		panic(err)
	}
	return r
}

// MustFastWithSource returns a fast derived Rand or panics.
func MustFastWithSource(src core.Source) Rand {
	r, err := FastWithSource(src)
	if err != nil {
		panic(err)
	}
	return r
}
