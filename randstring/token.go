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

// TokenHexBytes returns a lower-case hex token as a byte slice.
// Callers may zero the returned slice after use.
func TokenHexBytes(nBytes int) ([]byte, error) {
	return Default().TokenHexBytes(nBytes)
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

// TokenBase64Bytes returns a standard base64 token as a byte slice.
// Callers may zero the returned slice after use.
func TokenBase64Bytes(nBytes int) ([]byte, error) {
	return Default().TokenBase64Bytes(nBytes)
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

// TokenURLSafeBytes returns a URL-safe base64 token as a byte slice.
// Callers may zero the returned slice after use.
func TokenURLSafeBytes(nBytes int) ([]byte, error) {
	return Default().TokenURLSafeBytes(nBytes)
}
