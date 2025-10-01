package collection

import (
	"math"
	"sort"

	"github.com/aatuh/randutil/v2/core"
)

// Use core error sentinels for consistency across packages.

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
		return z, core.ErrEmptyItems
	}
	if len(items) != len(weights) {
		return z, core.ErrWeightsMismatch
	}
	var sum float64
	for _, w := range weights {
		if w < 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			return z, core.ErrInvalidWeights
		}
		sum += w
	}
	if sum <= 0 {
		return z, core.ErrInvalidWeights
	}
	u, err := Default.G.Float64()
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
	return z, core.ErrInvalidWeights
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
		return nil, core.ErrInvalidN
	}
	if k == 0 {
		return []T{}, nil
	}
	if len(items) == 0 {
		return nil, core.ErrEmptyItems
	}
	if len(items) != len(weights) {
		return nil, core.ErrWeightsMismatch
	}
	type kv struct {
		key float64
		i   int
	}
	keys := make([]kv, 0, len(items))
	for i, w := range weights {
		if w < 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			return nil, core.ErrInvalidWeights
		}
		if w == 0 {
			continue
		}
		// Draw u in (0,1], then key = -ln(u)/w. Smaller key = higher
		// chance. If u == 0, retry (extremely unlikely).
		var u float64
		for {
			var err error
			u, err = Default.G.Float64()
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
	if len(keys) == 0 {
		return nil, core.ErrInvalidWeights
	}
	if k > len(keys) {
		return nil, core.ErrSampleTooLarge
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
