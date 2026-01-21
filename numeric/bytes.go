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

// MustBytes returns n random bytes from crypto/rand. It panics on error.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: A slice of n random bytes.
func MustBytes(n int) []byte {
	b, err := Default().Bytes(n)
	if err != nil {
		panic(err)
	}
	return b
}

// MustBytes returns n random bytes from the generator's entropy source.
// It panics on error.
func (g *Generator) MustBytes(n int) []byte {
	b, err := g.Bytes(n)
	if err != nil {
		panic(err)
	}
	return b
}
