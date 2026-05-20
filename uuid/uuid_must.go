//go:build randutil_must
// +build randutil_must

package uuid

// MustV4 returns a v4 UUID or panics.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
func MustV4() UUID {
	u, err := V4()
	if err != nil {
		panic(err)
	}
	return u
}

// MustV7 returns a v7 UUID or panics.
//
// Returns:
//   - UUID: A random UUID conforming to Version 7 and Variant 1.
func MustV7() UUID {
	u, err := V7()
	if err != nil {
		panic(err)
	}
	return u
}

// MustParse panics on invalid input.
//
// Parameters:
//   - s: The string to parse.
//
// Returns:
//   - UUID: A lower-case UUID.
//   - error: An error if the string is invalid.
func MustParse(s string) UUID {
	u, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
