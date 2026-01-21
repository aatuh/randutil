package collection

// PickOne returns one random element from the slice.
// Returns an error if the slice is empty.
func PickOne[T any](slice []T) (T, error) {
	return Default[T]().PickOne(slice)
}

// MustPickOne returns one random element from the slice.
// It panics if an error occurs.
func MustPickOne[T any](slice []T) T {
	item, err := PickOne(slice)
	if err != nil {
		panic(err)
	}
	return item
}

// Shuffle performs an in-place secure Fisher-Yates shuffle of the slice.
func Shuffle[T any](slice []T) error {
	return Default[T]().Shuffle(slice)
}

// MustShuffle performs an in-place secure Fisher-Yates shuffle of the slice.
// It panics if an error occurs.
func MustShuffle[T any](slice []T) {
	if err := Shuffle(slice); err != nil {
		panic(err)
	}
}

// Choice returns a random choice from the provided arguments.
func Choice[T any](choices ...T) (T, error) {
	return PickOne(choices)
}

// MustChoice returns a random choice from the provided arguments.
// It panics if an error occurs.
func MustChoice[T any](choices ...T) T {
	return MustPickOne(choices)
}

// Sample returns k items uniformly at random from s without replacement.
// The input slice is not modified. If k == 0, an empty slice is returned.
// If k > len(s), ErrSampleTooLarge is returned.
func Sample[T any](s []T, k int) ([]T, error) {
	return Default[T]().Sample(s, k)
}

// MustSample returns k items uniformly at random from s without
// replacement. It panics on error.
func MustSample[T any](s []T, k int) []T {
	r, err := Sample(s, k)
	if err != nil {
		panic(err)
	}
	return r
}

// Perm returns a shuffled copy of slice.
func Perm[T any](slice []T) ([]T, error) {
	return Default[T]().Perm(slice)
}

// MustPerm returns a shuffled copy of slice. It panics on error.
func MustPerm[T any](slice []T) []T {
	p, err := Perm(slice)
	if err != nil {
		panic(err)
	}
	return p
}
