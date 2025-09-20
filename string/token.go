package string

import (
	"encoding/base64"
	"encoding/hex"
)

// TokenHex returns a lower-case hex string of length 2*nBytes.
// It reads nBytes from crypto/rand.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A lower-case hex string of length 2*nBytes.
//   - error: An error if crypto/rand fails.
func TokenHex(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
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

// TokenBase64 returns a standard base64 string (with padding) encoding
// nBytes of random data.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A standard base64 string (with padding) of length
//     approximately 4*ceil(nBytes/3).
//   - error: An error if crypto/rand fails.
func TokenBase64(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
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

// TokenURLSafe returns a URL-safe base64 string without padding
// encoding nBytes of random data.
//
// Parameters:
//   - nBytes: The length of the random bytes.
//
// Returns:
//   - string: A URL-safe base64 string without padding of length
//     approximately 4*ceil(nBytes/3).
//   - error: An error if crypto/rand fails.
func TokenURLSafe(nBytes int) (string, error) {
	b, err := Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
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
