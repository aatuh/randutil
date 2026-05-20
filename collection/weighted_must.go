//go:build randutil_must
// +build randutil_must

package collection

// MustWeightedChoice returns one item where probability is proportional to
// its non-negative weight. It panics on error.
func MustWeightedChoice[T any](items []T, weights []float64) T {
	item, err := WeightedChoice(items, weights)
	if err != nil {
		panic(err)
	}
	return item
}

// MustWeightedSample returns k distinct items without replacement, with
// probability proportional to weight. It panics on error.
func MustWeightedSample[T any](items []T, weights []float64, k int) []T {
	result, err := WeightedSample(items, weights, k)
	if err != nil {
		panic(err)
	}
	return result
}
