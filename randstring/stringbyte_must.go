//go:build randutil_must
// +build randutil_must

package randstring

// MustString returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the active entropy source.
// It panics on error.
//
// Parameters:
//   - length: The length of the string to generate.
//
// Returns:
//   - string: A random string of length characters drawn from the predefined
//     charset.
func MustString(length int) string {
	s, err := String(length)
	if err != nil {
		panic(err)
	}
	return s
}

// MustBase64 returns a base64 string built from byteLen random bytes. It panics on error.
//
// Parameters:
//   - byteLen: The number of random bytes to generate.
//
// Returns:
//   - string: A base64 string of length approximately 4*ceil(byteLen/3).
func MustBase64(byteLen int) string {
	s, err := Base64(byteLen)
	if err != nil {
		panic(err)
	}
	return s
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
