//go:build randutil_must
// +build randutil_must

package randstring

// MustTokenHex returns a lower-case hex string of length 2*nBytes.
// It panics on error.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A lower-case hex string of length 2*nBytes.
func MustTokenHex(nBytes int) string {
	s, err := TokenHex(nBytes)
	if err != nil {
		panic(err)
	}
	return s
}

// MustTokenHexBytes returns a lower-case hex token as a byte slice.
// It panics on error.
func MustTokenHexBytes(nBytes int) []byte {
	b, err := TokenHexBytes(nBytes)
	if err != nil {
		panic(err)
	}
	return b
}

// MustTokenBase64 returns a standard base64 string (with padding) encoding
// nBytes of random data. It panics on error.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A standard base64 string (with padding).
func MustTokenBase64(nBytes int) string {
	s, err := TokenBase64(nBytes)
	if err != nil {
		panic(err)
	}
	return s
}

// MustTokenBase64Bytes returns a standard base64 token as a byte slice.
// It panics on error.
func MustTokenBase64Bytes(nBytes int) []byte {
	b, err := TokenBase64Bytes(nBytes)
	if err != nil {
		panic(err)
	}
	return b
}

// MustTokenURLSafe returns a URL-safe base64 string without padding
// encoding nBytes of random data. It panics on error.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A URL-safe base64 string without padding.
func MustTokenURLSafe(nBytes int) string {
	s, err := TokenURLSafe(nBytes)
	if err != nil {
		panic(err)
	}
	return s
}

// MustTokenURLSafeBytes returns a URL-safe base64 token as a byte slice.
// It panics on error.
func MustTokenURLSafeBytes(nBytes int) []byte {
	b, err := TokenURLSafeBytes(nBytes)
	if err != nil {
		panic(err)
	}
	return b
}
