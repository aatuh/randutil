package collection

// PickByProbability returns items independently with probability p in [0,1].
// It preserves input order and allocates once.
func PickByProbability[T any](xs []T, p float64) ([]T, error) {
	return Default[T]().PickByProbability(xs, p)
}
