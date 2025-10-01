package numeric

// Bytes returns n random bytes from the active entropy source.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: A slice of n random bytes.
//   - error: An error if crypto/rand fails.
func Bytes(n int) ([]byte, error) {
	return def.Bytes(n)
}

// Fill populates b with random bytes from the active entropy source.
//
// Parameters:
//   - b: A byte slice to populate with random bytes.
//
// Returns:
//   - error: An error if crypto/rand fails.
func Fill(b []byte) error {
	return def.Fill(b)
}

// MustBytes returns n random bytes from crypto/rand. It panics on error.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: A slice of n random bytes.
func MustBytes(n int) []byte {
	b, err := Bytes(n)
	if err != nil {
		panic(err)
	}
	return b
}
