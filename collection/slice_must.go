//go:build randutil_must
// +build randutil_must

package collection

// MustPickOne returns one random element from the slice.
// It panics if an error occurs.
func MustPickOne[T any](slice []T) T {
	item, err := PickOne(slice)
	if err != nil {
		panic(err)
	}
	return item
}

// MustShuffle performs an in-place secure Fisher-Yates shuffle of the slice.
// It panics if an error occurs.
func MustShuffle[T any](slice []T) {
	if err := Shuffle(slice); err != nil {
		panic(err)
	}
}

// MustChoice returns a random choice from the provided arguments.
// It panics if an error occurs.
func MustChoice[T any](choices ...T) T {
	return MustPickOne(choices)
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

// MustPerm returns a shuffled copy of slice. It panics on error.
func MustPerm[T any](slice []T) []T {
	p, err := Perm(slice)
	if err != nil {
		panic(err)
	}
	return p
}
