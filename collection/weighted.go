package collection

// WeightedChoice returns one item where probability is proportional to
// its non-negative weight. Zero-weight items are never chosen.
func WeightedChoice[T any](items []T, weights []float64) (T, error) {
	return Default[T]().WeightedChoice(items, weights)
}

// WeightedSample returns k distinct items without replacement, with
// probability proportional to weight, using the Efraimidis–Spirakis
// exponential-keys method. Zero weights are never selected.
func WeightedSample[T any](
	items []T, weights []float64, k int,
) ([]T, error) {
	return Default[T]().WeightedSample(items, weights, k)
}
