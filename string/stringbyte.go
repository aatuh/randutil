package string

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"math/big"
)

// Predefined character sets.
const (
	lowerCase           = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	lowerCaseAndNumbers = lowerCase + numbers
)

// Uint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. n must be > 0.
func Uint64n(n uint64) (uint64, error) {
	if n == 0 {
		return 0, errors.New("n must be > 0")
	}
	var (
		max   = ^uint64(0)
		limit = max - (max % n)
	)
	for {
		var b [8]byte
		if _, err := rand.Read(b[:]); err != nil {
			return 0, err
		}
		u := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
			uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
		if u < limit {
			return u % n, nil
		}
	}
}

// Bytes returns n random bytes from crypto/rand.
func Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, errors.New("n must be non-negative")
	}
	if n == 0 {
		return []byte{}, nil
	}
	buf := make([]byte, n)
	if _, err := rand.Read(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// IntRange returns a secure random int in [minInclusive, maxInclusive].
func IntRange(minInclusive int, maxInclusive int) (int, error) {
	if minInclusive > maxInclusive {
		return 0, errors.New("min value is greater than max value")
	}
	diff := int64(maxInclusive) - int64(minInclusive) + 1
	rng := big.NewInt(diff)
	n, err := rand.Int(rand.Reader, rng)
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + minInclusive, nil
}

// GetWithCustomCharset returns a random string of length characters
// drawn from the provided charset using the active entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//   - charset: The charset to use for the string.
//
// Returns:
//   - string: A random string of length characters drawn from the provided charset.
//   - error: An error if crypto/rand fails.
func GetWithCustomCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", errors.New("length must be non-negative")
	}
	if len(charset) == 0 {
		return "", errors.New("charset must be non-empty")
	}
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		idx, err := Uint64n(uint64(len(charset)))
		if err != nil {
			return "", err
		}
		out[i] = charset[int(idx)]
	}
	return string(out), nil
}

// MustGetWithCustomCharset returns a random string of length characters
// drawn from the provided charset using the active entropy source. It panics
// on error.
//
// Parameters:
//   - length: The length of the string to generate.
//   - charset: The charset to use for the string.
//
// Returns:
//   - string: A random string of length characters drawn from the provided charset.
func MustGetWithCustomCharset(length int, charset string) string {
	s, err := GetWithCustomCharset(length, charset)
	if err != nil {
		panic(err)
	}
	return s
}

// String returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the active entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//
// Returns:
//   - string: A random string of length characters drawn from the predefined charset.
//   - error: An error if crypto/rand fails.
func String(length int) (string, error) {
	return GetWithCustomCharset(length, lowerCaseAndNumbers)
}

// MustString returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the active entropy source. It panics
// on error.
//
// Parameters:
//   - length: The length of the string to generate.
//
// Returns:
//   - string: A random string of length characters drawn from the predefined charset.
func MustString(length int) string {
	return MustGetWithCustomCharset(length, lowerCaseAndNumbers)
}

// Base64 returns a base64 string built from byteLen random bytes using
// the active entropy source.
//
// Parameters:
//   - byteLen: The length of the base64 string to generate.
//
// Returns:
//   - string: A base64 string of length byteLen.
//   - error: An error if crypto/rand fails.
func Base64(byteLen int) (string, error) {
	b, err := Bytes(byteLen)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// MustBase64 returns a base64 string of length byteLen. It panics on error.
//
// Parameters:
//   - byteLen: The length of the base64 string to generate.
//
// Returns:
//   - string: A base64 string of length byteLen.
func MustBase64(byteLen int) string {
	s, err := Base64(byteLen)
	if err != nil {
		panic(err)
	}
	return s
}

// Hex returns a lower-case hex string of length strLen. strLen must be
// even because each byte encodes to 2 hex chars.
//
// Parameters:
//   - strLen: The length of the hex string to generate.
//
// Returns:
//   - string: A lower-case hex string of length strLen.
//   - error: An error if crypto/rand fails.
func Hex(strLen int) (string, error) {
	if strLen%2 != 0 {
		return "", errors.New("strLen must be even")
	}
	b, err := Bytes(strLen / 2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// MustHex returns a lower-case hex string of length strLen. It panics on error.
//
// Parameters:
//   - strLen: The length of the hex string to generate.
//
// Returns:
//   - string: A lower-case hex string of length strLen.
func MustHex(strLen int) string {
	s, err := Hex(strLen)
	if err != nil {
		panic(err)
	}
	return s
}

// StringSlice returns a slice of random strings with per-item length in
// [minStrLen, maxStrLen], using the active entropy source.
//
// Parameters:
//   - sliceLength: The length of the slice to generate.
//   - minStrLen: The minimum length of the strings in the slice.
//   - maxStrLen: The maximum length of the strings in the slice.
//
// Returns:
//   - []string: A slice of random strings with per-item length in
//     [minStrLen, maxStrLen].
//   - error: An error if crypto/rand fails.
func StringSlice(sliceLength, minStrLen, maxStrLen int) ([]string, error) {
	result := make([]string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		sLen, err := IntRange(minStrLen, maxStrLen)
		if err != nil {
			return nil, err
		}
		s, err := String(sLen)
		if err != nil {
			return nil, err
		}
		result[i] = s
	}
	return result, nil
}

// MustStringSlice returns a slice of random strings with per-item length in
// [minStrLen, maxStrLen], using the active entropy source. It panics on error.
//
// Parameters:
//   - sliceLength: The length of the slice to generate.
//   - minStrLen: The minimum length of the strings in the slice.
//   - maxStrLen: The maximum length of the strings in the slice.
//
// Returns:
//   - []string: A slice of random strings with per-item length in
//     [minStrLen, maxStrLen].
func MustStringSlice(sliceLength, minStrLen, maxStrLen int) []string {
	s, err := StringSlice(sliceLength, minStrLen, maxStrLen)
	if err != nil {
		panic(err)
	}
	return s
}
