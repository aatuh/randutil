//go:build randutil_must
// +build randutil_must

package ulid

// MustULID returns a new ULID or panics on error.
func MustULID() ULID {
	u, err := ID()
	if err != nil {
		panic(err)
	}
	return u
}

// MustULID returns a new ULID or panics on error.
func (g *Generator) MustULID() ULID {
	u, err := g.ULID()
	if err != nil {
		panic(err)
	}
	return u
}

// MustParse panics on invalid input.
func MustParse(s string) ULID {
	u, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}
