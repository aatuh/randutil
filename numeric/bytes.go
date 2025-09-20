package numeric

import (
	"github.com/aatuh/randutil/core"
)

// Bytes returns n random bytes from the active entropy source.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: A slice of n random bytes.
//   - error: An error if crypto/rand fails.
func Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, core.ErrInvalidN
	}
	if n == 0 {
		return []byte{}, nil
	}
	buf := make([]byte, n)
	if err := readFull(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// Fill populates b with random bytes from the active entropy source.
//
// Parameters:
//   - b: A byte slice to populate with random bytes.
//
// Returns:
//   - error: An error if crypto/rand fails.
func Fill(b []byte) error {
	return readFull(b)
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
