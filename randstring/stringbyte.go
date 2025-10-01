package randstring

// StringWithCharset returns a random string of length characters
// drawn from the provided charset using the active entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//   - charset: The charset to use for the string.
//
// Returns:
//   - string: A random string of length characters drawn from the provided
//     charset.
//   - error: An error if crypto/rand fails.
func StringWithCharset(length int, charset string) (string, error) {
	return Default.StringWithCharset(length, charset)
}

// String returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the active entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//
// Returns:
//   - string: A random string of length characters drawn from the predefined
//     charset.
//   - error: An error if crypto/rand fails.
func String(length int) (string, error) {
	return Default.String(length)
}

// MustString returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the active entropy source.
// It panics
// on error.
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
	return Default.Base64(byteLen)
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
	return Default.Hex(strLen)
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
	return Default.StringSlice(sliceLength, minStrLen, maxStrLen)
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
