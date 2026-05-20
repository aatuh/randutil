//go:build randutil_must
// +build randutil_must

package nanoid

// MustID returns a NanoID with the default length. It panics on error.
func MustID() string {
	id, err := ID()
	if err != nil {
		panic(err)
	}
	return id
}

// MustIDWithLength returns a NanoID with the requested length. It panics on error.
func MustIDWithLength(length int) string {
	id, err := IDWithLength(length)
	if err != nil {
		panic(err)
	}
	return id
}

// MustID returns a NanoID with the requested length. It panics on error.
func (g *Generator) MustID(length int) string {
	id, err := g.ID(length)
	if err != nil {
		panic(err)
	}
	return id
}
