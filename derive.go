package randutil

import (
	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

// Derive returns a Rand bundle derived from seed and label.
// The derived stream is cryptographically strong if the seed is high entropy
// and kept secret. Deterministic seeds are for tests and replayable fixtures.
func Derive(seed []byte, label string) (Rand, error) {
	src, err := adapters.DeriveSource(seed, label)
	if err != nil {
		return Rand{}, err
	}
	return New(src), nil
}

// DeriveRNG returns a core RNG derived from seed and label.
func DeriveRNG(seed []byte, label string) (core.RNG, error) {
	return adapters.DeriveRNG(seed, label)
}
