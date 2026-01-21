package collection

// WeightedChoice returns one item where probability is proportional to
// its non-negative weight. Zero-weight items are never chosen.
func WeightedChoice[T any](items []T, weights []float64) (T, error) {
	return Default[T]().WeightedChoice(items, weights)
}

// MustWeightedChoice returns one item where probability is proportional to
// its non-negative weight. It panics on error.
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
func WeightedSample[T any](
	items []T, weights []float64, k int,
) ([]T, error) {
	return Default[T]().WeightedSample(items, weights, k)
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
