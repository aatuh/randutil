package randstring

import (
	"encoding/base64"
	"encoding/hex"
	"io"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds random strings/tokens using a source.
type Generator struct {
	G core.Generator
}

// New returns a string Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator { return &Generator{G: core.Generator{R: src}} }

// Default is the package-wide default generator.
var Default = New(nil)

// StringWithCharset returns a random string of length characters
// drawn from the provided charset using the generator's entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//   - charset: The character set to use for generation.
//
// Returns:
//   - string: A random string of the specified length.
//   - error: An error if length < 0, charset is empty, or if entropy fails.
func (g *Generator) StringWithCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", core.ErrInvalidN
	}
	if len(charset) == 0 {
		return "", core.ErrEmptyCharset
	}
	out := make([]byte, length)
	for i := 0; i < length; i++ {
		u, err := g.G.Uint64n(uint64(len(charset)))
		if err != nil {
			return "", err
		}
		out[i] = charset[int(u)]
	}
	return string(out), nil
}

// String returns a random string of length characters drawn from the
// predefined charset (lowercase+digits) using the generator's entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//
// Returns:
//   - string: A random string of the specified length.
//   - error: An error if length < 0 or if entropy fails.
func (g *Generator) String(length int) (string, error) {
	return g.StringWithCharset(length, lowerCaseAndNumbers)
}

// Base64 returns a base64 string built from byteLen random bytes using
// the generator's entropy source.
//
// Parameters:
//   - byteLen: The number of random bytes to generate.
//
// Returns:
//   - string: A base64-encoded string.
//   - error: An error if byteLen < 0 or if entropy fails.
func (g *Generator) Base64(byteLen int) (string, error) {
	b, err := g.G.Bytes(byteLen)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// Hex returns a lower-case hex string of length strLen. strLen must be
// even because each byte encodes to 2 hex chars.
//
// Parameters:
//   - strLen: The length of the hex string to generate.
//
// Returns:
//   - string: A lower-case hex string of length strLen.
//   - error: An error if strLen is odd, < 0, or if entropy fails.
func (g *Generator) Hex(strLen int) (string, error) {
	if strLen%2 != 0 {
		return "", core.ErrOddHexLength
	}
	b, err := g.G.Bytes(strLen / 2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TokenHex returns a lower-case hex string of length 2*nBytes.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A lower-case hex string of length 2*nBytes.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenHex(nBytes int) (string, error) {
	b, err := g.G.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TokenBase64 returns a standard base64 string (with padding) encoding
// nBytes of random data.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A standard base64 string with padding.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenBase64(nBytes int) (string, error) {
	b, err := g.G.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// TokenURLSafe returns a URL-safe base64 string without padding
// encoding nBytes of random data.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A URL-safe base64 string without padding.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenURLSafe(nBytes int) (string, error) {
	b, err := g.G.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// StringSlice returns a slice of random strings with per-item length in
// [minStrLen, maxStrLen], using the generator's entropy source.
func (g *Generator) StringSlice(
	sliceLength, minStrLen, maxStrLen int,
) ([]string, error) {
	if sliceLength < 0 {
		return nil, core.ErrInvalidN
	}
	result := make([]string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		sLen, err := g.G.IntRange(minStrLen, maxStrLen)
		if err != nil {
			return nil, err
		}
		s, err := g.String(sLen)
		if err != nil {
			return nil, err
		}
		result[i] = s
	}
	return result, nil
}

// Predefined character sets.
const (
	lowerCase           = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	lowerCaseAndNumbers = lowerCase + numbers
)
