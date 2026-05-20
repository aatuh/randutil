package collection

import (
	"math"
	"sort"

	"github.com/aatuh/randutil/v2/core"
)

func shuffleWithRNG[T any](rng rng, slice []T) error {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		bound, err := intToUint64(i + 1)
		if err != nil {
			return err
		}
		u, err := rng.Uint64n(bound)
		if err != nil {
			return err
		}
		j, err := uint64ToInt(u)
		if err != nil {
			return err
		}
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

func sampleWithRNG[T any](rng rng, s []T, k int) ([]T, error) {
	n := len(s)
	if k < 0 {
		return nil, core.ErrNegativeLength
	}
	if k == 0 {
		return []T{}, nil
	}
	if k > n {
		return nil, core.ErrSampleTooLarge
	}
	dup := make([]T, n)
	copy(dup, s)
	for i := 0; i < k; i++ {
		bound, err := intToUint64(n - i)
		if err != nil {
			return nil, err
		}
		u, err := rng.Uint64n(bound)
		if err != nil {
			return nil, err
		}
		j, err := uint64ToInt(u)
		if err != nil {
			return nil, err
		}
		j += i
		dup[i], dup[j] = dup[j], dup[i]
	}
	return dup[:k], nil
}

func pickOneWithRNG[T any](rng rng, slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, core.ErrEmptySlice
	}
	idx, err := rng.Intn(len(slice))
	if err != nil {
		var zero T
		return zero, err
	}
	return slice[idx], nil
}

func weightedChoiceWithRNG[T any](rng rng, items []T, weights []float64) (T, error) {
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
	if sum <= 0 || math.IsInf(sum, 0) || math.IsNaN(sum) {
		return z, core.ErrInvalidWeights
	}
	u, err := rng.Float64()
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
	for i := len(weights) - 1; i >= 0; i-- {
		if weights[i] > 0 {
			return items[i], nil
		}
	}
	return z, core.ErrInvalidWeights
}

func weightedSampleWithRNG[T any](
	rng rng, items []T, weights []float64, k int,
) ([]T, error) {
	if k < 0 {
		return nil, core.ErrNegativeLength
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
	if k > len(items) {
		return nil, core.ErrSampleTooLarge
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
		var u float64
		for {
			var err error
			u, err = rng.Float64()
			if err != nil {
				return nil, err
			}
			if u > 0 {
				break
			}
		}
		key := -math.Log(u) / w
		if key < 0 || math.IsNaN(key) || math.IsInf(key, 0) {
			return nil, core.ErrInvalidWeights
		}
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

func pickByProbabilityWithRNG[T any](rng rng, xs []T, p float64) ([]T, error) {
	if math.IsNaN(p) || math.IsInf(p, 0) || p < 0 || p > 1 {
		return nil, core.ErrInvalidProbability
	}
	if len(xs) == 0 {
		return []T{}, nil
	}
	if p == 0 {
		return []T{}, nil
	}
	if p == 1 {
		out := make([]T, len(xs))
		copy(out, xs)
		return out, nil
	}
	out := make([]T, 0, len(xs))
	for _, it := range xs {
		u, err := rng.Float64()
		if err != nil {
			return nil, err
		}
		if u < p {
			out = append(out, it)
		}
	}
	return out, nil
}
