package adapters

import (
	"io"

	"github.com/aatuh/randutil/v2/core"
)

const fastDeriveLabel = "randutil fast v1"

// FastSource returns a fast CSPRNG source derived from crypto/rand.
// For strict FIPS/OS RNG compliance, use crypto/rand.Reader directly.
func FastSource() (core.Source, error) {
	return FastSourceWithSource(nil)
}

// FastSourceWithSource returns a fast CSPRNG source derived from src.
// If src is nil, crypto/rand.Reader is used.
func FastSourceWithSource(src core.Source) (core.Source, error) {
	if src == nil {
		src = CryptoSource()
	}
	var seed [32]byte
	if _, err := io.ReadFull(src, seed[:]); err != nil {
		return nil, err
	}
	derived, err := DeriveSource(seed[:], fastDeriveLabel)
	core.Zero(seed[:])
	return derived, err
}

// FastRNG returns a fast CSPRNG RNG derived from crypto/rand.
func FastRNG() (core.RNG, error) {
	src, err := FastSource()
	if err != nil {
		return nil, err
	}
	return core.New(src), nil
}
