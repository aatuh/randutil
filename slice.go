package random

import "errors"

// SlicePickOne returns one random element from the slice.
// Returns an error if the slice is empty.
//
// Returns:
//   - any: A random element from the slice.
//   - error: An error if the slice is empty or if crypto/rand fails.
func SlicePickOne(slice []any) (any, error) {
	if len(slice) == 0 {
		var zero any
		return zero, errors.New("cannot pick from empty slice")
	}
	idx, err := IntRange(0, len(slice)-1)
	if err != nil {
		var zero any
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
func MustSlicePickOne(slice []any) any {
	item, err := SlicePickOne(slice)
	if err != nil {
		panic(err)
	}
	return item
}

// SlicePickMany returns a subset of items from the slice. For each item,
// a random chance is compared with chanceThreshold (0-100) to decide if it
// should be included.
//
// Parameters:
//   - slice: A slice of any type.
//   - chanceThreshold: An integer between 0 and 100.
//
// Returns:
//   - []any: A slice of any type.
//   - error: An error if crypto/rand fails.
func SlicePickMany(slice []any, chanceThreshold int) ([]any, error) {
	var picked []any
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
//   - []any: A slice of any type.
func MustSlicePickMany(slice []any, chanceThreshold int) []any {
	items, err := SlicePickMany(slice, chanceThreshold)
	if err != nil {
		panic(err)
	}
	return items
}

// Shuffle performs an in-place secure Fisher–Yates shuffle of the slice.
//
// Returns:
//   - error: An error if crypto/rand fails.
func Shuffle(slice []any) error {
	n := len(slice)
	for i := n - 1; i > 0; i-- {
		j, err := IntRange(0, i)
		if err != nil {
			return err
		}
		slice[i], slice[j] = slice[j], slice[i]
	}
	return nil
}

// MustShuffle performs an in-place secure Fisher–Yates shuffle of the slice.
// It panics if an error occurs.
//
// Parameters:
//   - slice: A slice of any type.
func MustShuffle(slice []any) {
	if err := Shuffle(slice); err != nil {
		panic(err)
	}
}

// Choice returns a random choice from the provided arguments.
//
// Returns:
//   - any: A random choice from the provided arguments.
//   - error: An error if crypto/rand fails.
func Choice(choices ...any) (any, error) {
	return SlicePickOne(choices)
}

// MustChoice returns a random choice from the provided arguments.
// It panics if an error occurs.
//
// Parameters:
//   - choices: A slice of any type.
func MustChoice(choices ...any) any {
	return MustSlicePickOne(choices)
}
