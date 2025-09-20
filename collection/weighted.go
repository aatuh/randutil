package collection

import (
	"crypto/rand"
	"errors"
	"math"
	"sort"
)

// Float64 returns a uniform random float64 in [0.0, 1.0) with 53 bits
// of precision built from crypto/rand.
func Float64() (float64, error) {
	var b [8]byte
	if _, err := rand.Read(b[:]); err != nil {
		return 0, err
	}
	u := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	u >>= 11
	const denom = 1 << 53
	return float64(u) / float64(denom), nil
}

// Package-level errors.
var (
	ErrInvalidWeights  = errors.New("invalid weights")
	ErrWeightsMismatch = errors.New("items/weights mismatch")
	ErrInvalidN        = errors.New("n must be > 0")
	ErrSampleTooLarge  = errors.New("sample k exceeds size")
)

// WeightedChoice returns one item where probability is proportional to
// its non-negative weight. Zero-weight items are never chosen.
//
// Parameters:
//   - items:    Items to choose from.
//   - weights:  Non-negative weights, same length as items.
//
// Returns:
//   - T:    A chosen item.
//   - error: If inputs are invalid or entropy fails.
func WeightedChoice[T any](items []T, weights []float64) (T, error) {
	var z T
	if len(items) == 0 {
		return z, errors.New("randutil: empty items")
	}
	if len(items) != len(weights) {
		return z, ErrWeightsMismatch
	}
	var sum float64
	for _, w := range weights {
		if w < 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			return z, ErrInvalidWeights
		}
		sum += w
	}
	if sum <= 0 {
		return z, ErrInvalidWeights
	}
	u, err := Float64()
	if err != nil {
		return z, err
	}
	target := u * sum
	var acc float64
	for i, w := range weights {
		acc += w
		if target < acc {
			return items[i], nil
		}
	}
	// Fallback (floating point corner), pick last positive weight.
	for i := len(weights) - 1; i >= 0; i-- {
		if weights[i] > 0 {
			return items[i], nil
		}
	}
	return z, ErrInvalidWeights
}

// MustWeightedChoice returns one item where probability is proportional to
// its non-negative weight. It panics on error.
//
// Parameters:
//   - items:    Items to choose from.
//   - weights:  Non-negative weights, same length as items.
//
// Returns:
//   - T:    A chosen item.
func MustWeightedChoice[T any](items []T, weights []float64) T {
	item, err := WeightedChoice(items, weights)
	if err != nil {
		panic(err)
	}
	return item
}

// WeightedSample returns k distinct items without replacement, with
// probability proportional to weight, using the Efraimidis–Spirakis
// exponential-keys method. Zero weights are never selected.
//
// Parameters:
//   - items:    Items to sample from.
//   - weights:  Non-negative weights, same length as items.
//   - k:        Sample size.
//
// Returns:
//   - []T:  k sampled items.
//   - error: If inputs are invalid or entropy fails.
func WeightedSample[T any](
	items []T, weights []float64, k int,
) ([]T, error) {
	if k < 0 {
		return nil, errors.New("n must be > 0")
	}
	if k == 0 {
		return []T{}, nil
	}
	if len(items) == 0 {
		return nil, errors.New("randutil: empty items")
	}
	if len(items) != len(weights) {
		return nil, ErrWeightsMismatch
	}
	type kv struct {
		key float64
		i   int
	}
	keys := make([]kv, 0, len(items))
	for i, w := range weights {
		if w < 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			return nil, ErrInvalidWeights
		}
		if w == 0 {
			continue
		}
		// Draw u in (0,1], then key = -ln(u)/w. Smaller key = higher
		// chance. If u == 0, retry (extremely unlikely).
		var u float64
		for {
			var err error
			u, err = Float64()
			if err != nil {
				return nil, err
			}
			if u > 0 {
				break
			}
		}
		key := -math.Log(u) / w
		keys = append(keys, kv{key: key, i: i})
	}
	if k > len(keys) {
		return nil, errors.New("sample k exceeds size")
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i].key < keys[j].key })
	out := make([]T, k)
	for j := 0; j < k; j++ {
		out[j] = items[keys[j].i]
	}
	return out, nil
}

// MustWeightedSample returns k distinct items without replacement, with
// probability proportional to weight. It panics on error.
//
// Parameters:
//   - items:    Items to sample from.
//   - weights:  Non-negative weights, same length as items.
//   - k:        Sample size.
//
// Returns:
//   - []T:  k sampled items.
func MustWeightedSample[T any](items []T, weights []float64, k int) []T {
	result, err := WeightedSample(items, weights, k)
	if err != nil {
		panic(err)
	}
	return result
}
