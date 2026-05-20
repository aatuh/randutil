package numeric

// Bytes returns n random bytes from the default entropy source.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: A slice of n random bytes.
//   - error: An error if crypto/rand fails.
func Bytes(n int) ([]byte, error) { return Default().Bytes(n) }

// Bytes returns n random bytes from the generator's entropy source.
func (g *Generator) Bytes(n int) ([]byte, error) {
	return g.rng.Bytes(n)
}

// Fill populates b with random bytes from the active entropy source.
//
// Parameters:
//   - b: A byte slice to populate with random bytes.
//
// Returns:
//   - error: An error if crypto/rand fails.
func Fill(b []byte) error { return Default().Fill(b) }

// Fill populates b with random bytes from the generator's entropy source.
func (g *Generator) Fill(b []byte) error {
	return g.rng.Fill(b)
}
