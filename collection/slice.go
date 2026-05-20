package collection

// PickOne returns one random element from the slice.
// Returns an error if the slice is empty.
func PickOne[T any](slice []T) (T, error) {
	return Default[T]().PickOne(slice)
}

// Shuffle performs an in-place secure Fisher-Yates shuffle of the slice.
func Shuffle[T any](slice []T) error {
	return Default[T]().Shuffle(slice)
}

// Choice returns a random choice from the provided arguments.
func Choice[T any](choices ...T) (T, error) {
	return PickOne(choices)
}

// Sample returns k items uniformly at random from s without replacement.
// The input slice is not modified. If k == 0, an empty slice is returned.
// If k > len(s), ErrSampleTooLarge is returned.
func Sample[T any](s []T, k int) ([]T, error) {
	return Default[T]().Sample(s, k)
}

// Perm returns a shuffled copy of slice.
func Perm[T any](slice []T) ([]T, error) {
	return Default[T]().Perm(slice)
}
