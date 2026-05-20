package randutil

import (
	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

// Fast returns a Rand backed by a fast derived CSPRNG seeded from crypto/rand.
// For strict FIPS/OS RNG compliance, use Secure/Default instead.
func Fast() (Rand, error) {
	return FastWithSource(nil)
}

// FastWithSource returns a Rand backed by a fast derived CSPRNG seeded from src.
// If src is nil, crypto/rand.Reader is used.
func FastWithSource(src core.Source) (Rand, error) {
	derived, err := adapters.FastSourceWithSource(src)
	if err != nil {
		return Rand{}, err
	}
	return New(derived), nil
}
