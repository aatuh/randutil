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
