package collection

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Intn returns a uniform random int in [0, n). n must be > 0.
func Intn(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be > 0")
	}
	u, err := Uint64n(uint64(n))
	if err != nil {
		return 0, err
	}
	return int(u), nil
}

// Uint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. n must be > 0.
func Uint64n(n uint64) (uint64, error) {
	if n == 0 {
		return 0, errors.New("n must be > 0")
	}
	var (
		max   = ^uint64(0)
		limit = max - (max % n)
	)
	for {
		var b [8]byte
		if _, err := rand.Read(b[:]); err != nil {
			return 0, err
		}
		u := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
			uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
		if u < limit {
			return u % n, nil
		}
	}
}

// IntRange returns a secure random int in [minInclusive, maxInclusive].
func IntRange(minInclusive int, maxInclusive int) (int, error) {
	if minInclusive > maxInclusive {
		return 0, errors.New("min value is greater than max value")
	}
	diff := int64(maxInclusive) - int64(minInclusive) + 1
	rng := big.NewInt(diff)
	n, err := rand.Int(rand.Reader, rng)
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + minInclusive, nil
}

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
		return zero, errors.New("cannot pick from empty slice")
	}
	idx, err := Intn(len(slice))
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

// SlicePickMany returns a subset of items from the slice. For each item,
// a random chance is compared with chanceThreshold (0-100) to decide if
// it should be included.
//
// Parameters:
//   - slice: A slice of any type.
//   - chanceThreshold: An integer between 0 and 100.
//
// Returns:
//   - []T: A slice of the same type as input.
//   - error: An error if entropy fails.
func SlicePickMany[T any](slice []T, chanceThreshold int) ([]T, error) {
	var picked []T
	for _, item := range slice {
		dice, err := IntRange(0, 100)
		if err != nil {
			return nil, err
		}
		if dice <= chanceThreshold {
			picked = append(picked, item)
		}
	}
	return picked, nil
}

// MustSlicePickMany returns a subset of items from the slice.
// It panics if an error occurs.
//
// Parameters:
//   - slice: A slice of any type.
//   - chanceThreshold: An integer between 0 and 100.
//
// Returns:
//   - []T: A slice of the same type as input.
func MustSlicePickMany[T any](slice []T, chanceThreshold int) []T {
	items, err := SlicePickMany(slice, chanceThreshold)
	if err != nil {
		panic(err)
	}
	return items
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
		u, err := Uint64n(uint64(i + 1))
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
		return nil, errors.New("n must be > 0")
	}
	if k == 0 {
		return []T{}, nil
	}
	if k > n {
		return nil, errors.New("sample k exceeds size")
	}
	dup := make([]T, n)
	copy(dup, s)
	for i := 0; i < k; i++ {
		u, err := Uint64n(uint64(n - i))
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
		return nil, errors.New("n must be > 0")
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

// MustPerm returns a random permutation of the integers [0..n). It panics on error.
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
