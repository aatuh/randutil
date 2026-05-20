//go:build randutil_must
// +build randutil_must

package collection

// MustPickByProbability returns items independently with probability p in [0,1].
// It panics on error.
func MustPickByProbability[T any](xs []T, p float64) []T {
	result, err := PickByProbability(xs, p)
	if err != nil {
		panic(err)
	}
	return result
}
