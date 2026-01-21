package numeric

// Bool returns a secure random boolean.
//
// Returns:
//   - bool: A random boolean value.
//   - error
func Bool() (bool, error) { return Default().Bool() }

// Bool returns a secure random boolean from the generator's entropy source.
func (g *Generator) Bool() (bool, error) {
	return g.rng.Bool()
}

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
