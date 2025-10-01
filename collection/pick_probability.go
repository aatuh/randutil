package collection

import (
	"github.com/aatuh/randutil/v2/core"
)

// PickByProbability returns items independently with probability p in [0,1].
// It preserves input order and allocates once.
//
// Parameters:
//   - xs: The slice of items to pick from.
//   - p: The probability of selecting each item (must be in [0,1]).
//
// Returns:
//   - []T: A slice of selected items (preserves input order).
//   - error: An error if p is out of range or if entropy fails.
func PickByProbability[T any](xs []T, p float64) ([]T, error) {
	if p < 0 || p > 1 {
		return nil, core.ErrInvalidRange
	}
	if len(xs) == 0 {
		return []T{}, nil
	}
	out := make([]T, 0, len(xs)) // upper bound
	for _, it := range xs {
		u, err := Default.G.Float64()
		if err != nil {
			return nil, err
		}
		if u <= p {
			out = append(out, it)
		}
	}
	return out, nil
}

// MustPickByProbability returns items independently with probability p in [0,1].
// It panics on error.
//
// Parameters:
//   - xs: The slice of items to pick from.
//   - p: The probability of selecting each item (must be in [0,1]).
//
// Returns:
//   - []T: A slice of selected items (preserves input order).
func MustPickByProbability[T any](xs []T, p float64) []T {
	result, err := PickByProbability(xs, p)
	if err != nil {
		panic(err)
	}
	return result
}
