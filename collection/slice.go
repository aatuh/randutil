package collection

import (
	"github.com/aatuh/randutil/v2/core"
)

// SlicePickOne returns one random element from the slice.
// Returns an error if the slice is empty.
//
// Parameters:
//   - slice: A slice of any type.
//
// Returns:
//   - any: A random element from the slice.
//   - error: An error if the slice is empty or if entropy fails.
func SlicePickOne[T any](slice []T) (T, error) {
	if len(slice) == 0 {
		var zero T
		return zero, core.ErrEmptySlice
	}
	idx, err := Default.G.Intn(len(slice))
	if err != nil {
		var zero T
		return zero, err
	}
	return slice[idx], nil
}

// MustSlicePickOne returns one random element from the slice.
// It panics if an error occurs.
//
// Parameters:
//   - slice: A slice of any type.
//
// Returns:
//   - any: A random element from the slice.
func MustSlicePickOne[T any](slice []T) T {
	item, err := SlicePickOne(slice)
	if err != nil {
		panic(err)
	}
	return item
}

// Shuffle performs an in-place secure Fisher-Yates shuffle of the slice.
//
// Parameters:
//   - slice: A slice of any type to shuffle in-place.
//
// Returns:
//   - error: An error if entropy fails.
func Shuffle[T any](slice []T) error {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		u, err := Default.G.Uint64n(uint64(i + 1))
		if err != nil {
			return err
		}
		j := int(u)
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

// MustShuffle performs an in-place secure Fisher-Yates shuffle of the slice.
// It panics if an error occurs.
//
// Parameters:
//   - slice: A slice of any type.
func MustShuffle[T any](slice []T) {
	if err := Shuffle(slice); err != nil {
		panic(err)
	}
}

// Choice returns a random choice from the provided arguments.
//
// Parameters:
//   - choices: Variable number of arguments of any type.
//
// Returns:
//   - T: A random choice from the provided arguments.
//   - error: An error if entropy fails.
func Choice[T any](choices ...T) (T, error) {
	return SlicePickOne(choices)
}

// MustChoice returns a random choice from the provided arguments.
// It panics if an error occurs.
//
// Parameters:
//   - choices: Variable number of arguments of any type.
//
// Returns:
//   - T: A random choice from the provided arguments.
func MustChoice[T any](choices ...T) T {
	return MustSlicePickOne(choices)
}

// Sample returns k items uniformly at random from s without replacement.
// The input slice is not modified. If k == 0, an empty slice is returned.
// If k > len(s), ErrSampleTooLarge is returned.
//
// Implementation does a partial Fisher-Yates on a copy, swapping only
// the first k positions. For very large n and small k, a Floyd algo
// could be added later.
//
// Parameters:
//   - s: A slice of any type to sample from.
//   - k: The number of items to sample.
//
// Returns:
//   - []T: A slice of k items sampled from s.
//   - error: An error if entropy fails or k > len(s).
func Sample[T any](s []T, k int) ([]T, error) {
	n := len(s)
	if k < 0 {
		return nil, core.ErrInvalidN
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
		u, err := Default.G.Uint64n(uint64(n - i))
		if err != nil {
			return nil, err
		}
		j := int(u) + i
		dup[i], dup[j] = dup[j], dup[i]
	}
	return dup[:k], nil
}

// MustSample returns k items uniformly at random from s without
// replacement. It panics on error.
//
// Parameters:
//   - s: A slice of any type to sample from.
//   - k: The number of items to sample.
//
// Returns:
//   - []T: A slice of k items sampled from s.
func MustSample[T any](s []T, k int) []T {
	r, err := Sample(s, k)
	if err != nil {
		panic(err)
	}
	return r
}

// Perm returns a random permutation of the integers [0..n).
//
// Parameters:
//   - n: The upper bound (exclusive) for the permutation.
//
// Returns:
//   - []int: A random permutation of integers [0..n).
//   - error: An error if entropy fails or n < 0.
func Perm(n int) ([]int, error) {
	if n < 0 {
		return nil, core.ErrInvalidN
	}
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	if err := Shuffle(p); err != nil {
		return nil, err
	}
	return p, nil
}

// MustPerm returns a random permutation of the integers [0..n). It panics on
// error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the permutation.
//
// Returns:
//   - []int: A random permutation of integers [0..n).
func MustPerm(n int) []int {
	p, err := Perm(n)
	if err != nil {
		panic(err)
	}
	return p
}
