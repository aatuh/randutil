//go:build randutil_must
// +build randutil_must

package numeric

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
