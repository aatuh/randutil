//go:build randutil_must
// +build randutil_must

package numeric

// MustBool returns a secure random boolean. It panics if an error occurs.
//
// Returns:
//   - bool: A random boolean value.
func MustBool() bool {
	b, err := Default().Bool()
	if err != nil {
		panic(err)
	}
	return b
}

// MustBool returns a secure random boolean. It panics on error.
func (g *Generator) MustBool() bool {
	b, err := g.Bool()
	if err != nil {
		panic(err)
	}
	return b
}
