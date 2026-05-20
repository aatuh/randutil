package deterministic

import (
	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

// Source returns a reproducible stream based on seed.
func Source(seed []byte) (core.Source, error) {
	return adapters.DeterministicSource(seed)
}

// SourceWithLabel returns a reproducible stream based on seed and label.
func SourceWithLabel(seed []byte, label string) (core.Source, error) {
	return adapters.DeterministicSourceWithLabel(seed, label)
}

// RNG returns a core RNG derived from seed.
func RNG(seed []byte) (core.RNG, error) {
	src, err := Source(seed)
	if err != nil {
		return nil, err
	}
	return core.New(src), nil
}

// RNGWithLabel returns a core RNG derived from seed and label.
func RNGWithLabel(seed []byte, label string) (core.RNG, error) {
	src, err := SourceWithLabel(seed, label)
	if err != nil {
		return nil, err
	}
	return core.New(src), nil
}
