package collection

// PickByProbability returns items independently with probability p in [0,1].
// It preserves input order and allocates once.
func PickByProbability[T any](xs []T, p float64) ([]T, error) {
	return Default[T]().PickByProbability(xs, p)
}

// MustPickByProbability returns items independently with probability p in [0,1].
// It panics on error.
func MustPickByProbability[T any](xs []T, p float64) []T {
	result, err := PickByProbability(xs, p)
	if err != nil {
		panic(err)
	}
	return result
}
