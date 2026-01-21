package randstring

// TokenHex returns a lower-case hex string of length 2*nBytes.
// It reads nBytes from crypto/rand.
// Note: strings are immutable; use TokenHexBytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A lower-case hex string of length 2*nBytes.
//   - error: An error if crypto/rand fails.
func TokenHex(nBytes int) (string, error) {
	return Default().TokenHex(nBytes)
}

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

// TokenHexBytes returns a lower-case hex token as a byte slice.
// Callers may zero the returned slice after use.
func TokenHexBytes(nBytes int) ([]byte, error) {
	return Default().TokenHexBytes(nBytes)
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

// TokenBase64 returns a standard base64 string (with padding) encoding
// nBytes of random data.
// Note: strings are immutable; use TokenBase64Bytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A standard base64 string (with padding) of length
//     approximately 4*ceil(nBytes/3).
//   - error: An error if crypto/rand fails.
func TokenBase64(nBytes int) (string, error) {
	return Default().TokenBase64(nBytes)
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

// TokenBase64Bytes returns a standard base64 token as a byte slice.
// Callers may zero the returned slice after use.
func TokenBase64Bytes(nBytes int) ([]byte, error) {
	return Default().TokenBase64Bytes(nBytes)
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

// TokenURLSafe returns a URL-safe base64 string without padding
// encoding nBytes of random data.
// Note: strings are immutable; use TokenURLSafeBytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A URL-safe base64 string without padding of length
//     approximately 4*ceil(nBytes/3).
//   - error: An error if crypto/rand fails.
func TokenURLSafe(nBytes int) (string, error) {
	return Default().TokenURLSafe(nBytes)
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

// TokenURLSafeBytes returns a URL-safe base64 token as a byte slice.
// Callers may zero the returned slice after use.
func TokenURLSafeBytes(nBytes int) ([]byte, error) {
	return Default().TokenURLSafeBytes(nBytes)
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
