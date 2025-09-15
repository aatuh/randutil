package random

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"math/big"
)

// Predefined character sets.
const (
	lowerCase           = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	lowerCaseAndNumbers = lowerCase + numbers
)

// GetWithCustomCharset returns a random string of length characters drawn from
// the provided charset.
//
// Parameters:
//   - length: The length of the random string.
//   - charset: The characters to draw from.
//
// Returns:
//   - string: A random string of length characters drawn from the provided
//     charset.
//   - error: An error if crypto/rand fails.
func GetWithCustomCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", errors.New("length must be non-negative")
	}
	result := make([]byte, length)
	charsetLen := big.NewInt(int64(len(charset)))
	for i := 0; i < length; i++ {
		n, err := rand.Int(rand.Reader, charsetLen)
		if err != nil {
			return "", err
		}
		result[i] = charset[n.Int64()]
	}
	return string(result), nil
}

// MustGetWithCustomCharset returns a random string of length characters drawn
// from the provided charset. It panics if an error occurs.
func MustGetWithCustomCharset(length int,
	charset string) string {
	s, err := GetWithCustomCharset(length, charset)
	if err != nil {
		panic(err)
	}
	return s
}

// String returns a secure random string of the given length using the
// default charset (lowerCaseAndNumbers).
//
// Parameters:
//   - length: The length of the random string.
//
// Returns:
//   - string: A random string of length characters drawn from the default
//     charset.
//   - error: An error if crypto/rand fails.
func String(length int) (string, error) {
	return GetWithCustomCharset(length, lowerCaseAndNumbers)
}

// MustString returns a secure random string of the given length using the
// default charset (lowerCaseAndNumbers). It panics if an error occurs.
//
// Parameters:
//   - length: The length of the random string.
//
// Returns:
//   - string: A random string of length characters drawn from the default
func MustString(length int) string {
	return MustGetWithCustomCharset(length, lowerCaseAndNumbers)
}

// Base64 returns a base64 encoded random string built from byteLen random
// bytes.
//
// Parameters:
//   - byteLen: The length of the random bytes.
//
// Returns:
//   - string: Base64 encoded random string.
//   - error: An error if crypto/rand fails.
func Base64(byteLen int) (string, error) {
	bytes := make([]byte, byteLen)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// Base64 returns a base64 encoded random string built from byteLen random
// bytes. It panics if an error occurs.
//
// Parameters:
//   - byteLen: The length of the random bytes.
//
// Returns:
//   - string: Base64 encoded random string.
func MustBase64(byteLen int) string {
	s, err := Base64(byteLen)
	if err != nil {
		panic(err)
	}
	return s
}

// Hex returns a random hexadecimal string of length strLen. strLen must be
// even or it will return an error.
//
// Parameters:
//   - strLen: The length of the random hexadecimal string.
//
// Returns:
//   - string: A random hexadecimal string of length strLen.
//   - error: An error if crypto/rand fails or strLen is not even.
func Hex(strLen int) (string, error) {
	if strLen%2 != 0 {
		return "", errors.New("strLen must be even")
	}
	bytes := make([]byte, strLen/2)
	if _, err := io.ReadFull(rand.Reader, bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", bytes), nil
}

// MustHex returns a random hexadecimal string of length strLen.
// It panics if an error occurs.
//
// Parameters:
//   - strLen: The length of the random hexadecimal string.
//
// Returns:
//   - string: A random hexadecimal string of length strLen.
func MustHex(strLen int) string {
	s, err := Hex(strLen)
	if err != nil {
		panic(err)
	}
	return s
}

// StringSlice returns a slice of random strings. Each string length is randomly
// chosen between minStrLen and maxStrLen.
//
// Parameters:
//   - sliceLength: The length of the slice.
//   - minStrLen: The minimum length of each string.
//   - maxStrLen: The maximum length of each string.
//
// Returns:
//   - []string: A slice of random strings.
//   - error: An error if crypto/rand fails.
func StringSlice(
	sliceLength int, minStrLen int, maxStrLen int,
) ([]string, error) {
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

// StringSlice returns a slice of random strings. It panics if an error occurs.
//
// Parameters:
//   - sliceLength: The length of the slice.
//   - minStrLen: The minimum length of each string.
//   - maxStrLen: The maximum length of each string.
//
// Returns:
//   - []string: A slice of random strings.
func MustStringSlice(sliceLength, minStrLen,
	maxStrLen int) []string {
	s, err := StringSlice(sliceLength, minStrLen, maxStrLen)
	if err != nil {
		panic(err)
	}
	return s
}
