package collection

import (
	"io"
	"math"
	"sort"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds slice-related random operations using a core generator.
type Generator struct {
	G core.Generator
}

// New returns a collection Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator { return &Generator{G: core.Generator{R: src}} }

// Default is the package-wide default generator.
var Default = New(nil)

// Shuffle performs an in-place secure Fisher-Yates shuffle of the slice.
//
// Parameters:
//   - slice: The slice to shuffle in-place.
//
// Returns:
//   - error: An error if entropy fails.
func (g *Generator) Shuffle(slice []int) error {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		u, err := g.G.Uint64n(uint64(i + 1))
		if err != nil {
			return err
		}
		j := int(u)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

// Sample returns k items uniformly at random from s without replacement.
// The input slice is not modified. If k == 0, an empty slice is returned.
// If k > len(s), ErrSampleTooLarge is returned.
//
// Parameters:
//   - s: The slice to sample from.
//   - k: The number of items to sample.
//
// Returns:
//   - []int: A slice of k items sampled from s.
//   - error: An error if k < 0, k > len(s), or if entropy fails.
func (g *Generator) Sample(s []int, k int) ([]int, error) {
	n := len(s)
	if k < 0 {
		return nil, core.ErrInvalidN
	}
	if k == 0 {
		return []int{}, nil
	}
	if k > n {
		return nil, core.ErrSampleTooLarge
	}
	dup := make([]int, n)
	copy(dup, s)
	for i := 0; i < k; i++ {
		u, err := g.G.Uint64n(uint64(n - i))
		if err != nil {
			return nil, err
		}
		j := int(u) + i
		dup[i], dup[j] = dup[j], dup[i]
	}
	return dup[:k], nil
}

// PickOne returns one random element from the slice.
// Returns an error if the slice is empty.
//
// Parameters:
//   - slice: The slice to pick from.
//
// Returns:
//   - int: A random element from the slice.
//   - error: An error if the slice is empty or if entropy fails.
func (g *Generator) PickOne(slice []int) (int, error) {
	if len(slice) == 0 {
		return 0, core.ErrEmptySlice
	}
	idx, err := g.G.Intn(len(slice))
	if err != nil {
		return 0, err
	}
	return slice[idx], nil
}

// Perm returns a random permutation of the integers [0..n).
//
// Parameters:
//   - n: The upper bound (exclusive) for the permutation.
//
// Returns:
//   - []int: A random permutation of integers [0..n).
//   - error: An error if n < 0 or if entropy fails.
func (g *Generator) Perm(n int) ([]int, error) {
	if n < 0 {
		return nil, core.ErrInvalidN
	}
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	if err := g.Shuffle(p); err != nil {
		return nil, err
	}
	return p, nil
}

// WeightedChoice returns one item where probability is proportional to
// its non-negative weight. Zero-weight items are never chosen.
//
// Parameters:
//   - items: The items to choose from.
//   - weights: The weights for each item (must be non-negative).
//
// Returns:
//   - int: A chosen item.
//   - error: An error if inputs are invalid or if entropy fails.
func (g *Generator) WeightedChoice(items []int, weights []float64) (int, error) {
	if len(items) == 0 {
		return 0, core.ErrEmptyItems
	}
	if len(items) != len(weights) {
		return 0, core.ErrWeightsMismatch
	}
	var sum float64
	for _, w := range weights {
		if w < 0 || math.IsNaN(w) || math.IsInf(w, 0) {
			return 0, core.ErrInvalidWeights
		}
		sum += w
	}
	if sum <= 0 {
		return 0, core.ErrInvalidWeights
	}
	u, err := g.G.Float64()
	if err != nil {
		return 0, err
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
	return 0, core.ErrInvalidWeights
}

// WeightedSample returns k items where probability is proportional to
// their non-negative weights. Zero-weight items are never chosen.
//
// Parameters:
//   - items: The items to sample from.
//   - weights: The weights for each item (must be non-negative).
//   - k: The number of items to sample.
//
// Returns:
//   - []int: k sampled items.
//   - error: An error if inputs are invalid or if entropy fails.
func (g *Generator) WeightedSample(items []int, weights []float64, k int) ([]int, error) {
	if k < 0 {
		return nil, core.ErrInvalidN
	}
	if k == 0 {
		return []int{}, nil
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
		// Draw u in (0,1], then key = -ln(u)/w. Smaller key = higher
		// chance. If u == 0, retry (extremely unlikely).
		var u float64
		for {
			var err error
			u, err = g.G.Float64()
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
	out := make([]int, k)
	for j := 0; j < k; j++ {
		out[j] = items[keys[j].i]
	}
	return out, nil
}
