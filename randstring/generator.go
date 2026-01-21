package randstring

import (
	"encoding/base64"
	"encoding/hex"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds random strings/tokens using a source.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng core.RNG
}

// New returns a string Generator. If rng is nil, crypto/rand is used.
func New(rng core.RNG) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator{rng: rng}
}

// NewWithSource returns a string Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return New(core.New(src))
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// StringWithCharset returns a random string of length characters
// drawn from the provided charset using the generator's entropy source.
//
// Parameters:
//   - length: The length of the string to generate.
//   - charset: The character set to use for generation.
//
// Returns:
//   - string: A random string of the specified length.
//   - error: An error if length < 0, charset is empty/invalid, or if entropy fails.
func (g *Generator) StringWithCharset(length int, charset string) (string, error) {
	if length < 0 {
		return "", core.ErrNegativeLength
	}
	if len(charset) == 0 {
		return "", core.ErrEmptyCharset
	}
	if !isASCIICharset(charset) {
		return "", core.ErrInvalidCharset
	}
	if length == 0 {
		return "", nil
	}
	charsetBytes := []byte(charset)
	if len(charsetBytes) == 1 {
		out := make([]byte, length)
		for i := range out {
			out[i] = charsetBytes[0]
		}
		return string(out), nil
	}

	out := make([]byte, length)
	n := len(charsetBytes)
	if n > 256 {
		for i := range out {
			idx, err := g.rng.Uint64n(uint64(n))
			if err != nil {
				return "", err
			}
			pos, err := u64ToInt(idx)
			if err != nil {
				return "", err
			}
			out[i] = charsetBytes[pos]
		}
		return string(out), nil
	}

	buf := make([]byte, 128)
	if isPowerOfTwo(n) {
		mask := byte(n - 1)
		pos := 0
		for pos < length {
			if err := g.rng.Fill(buf); err != nil {
				return "", err
			}
			for _, b := range buf {
				out[pos] = charsetBytes[int(b&mask)]
				pos++
				if pos == length {
					break
				}
			}
		}
		return string(out), nil
	}

	acceptLimit := byte(256 - (256 % n))
	pos := 0
	for pos < length {
		if err := g.rng.Fill(buf); err != nil {
			return "", err
		}
		for _, b := range buf {
			if b >= acceptLimit {
				continue
			}
			out[pos] = charsetBytes[int(b)%n]
			pos++
			if pos == length {
				break
			}
		}
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
	b, err := g.rng.Bytes(byteLen)
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
	b, err := g.rng.Bytes(strLen / 2)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TokenHex returns a lower-case hex string of length 2*nBytes.
// Note: strings are immutable; use TokenHexBytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A lower-case hex string of length 2*nBytes.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenHex(nBytes int) (string, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// TokenHexBytes returns a lower-case hex token as a byte slice.
// Callers may zero the returned slice after use.
func (g *Generator) TokenHexBytes(nBytes int) ([]byte, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return nil, err
	}
	out := make([]byte, hex.EncodedLen(len(b)))
	hex.Encode(out, b)
	return out, nil
}

// TokenBase64 returns a standard base64 string (with padding) encoding
// nBytes of random data.
// Note: strings are immutable; use TokenBase64Bytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A standard base64 string with padding.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenBase64(nBytes int) (string, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

// TokenBase64Bytes returns a standard base64 token as a byte slice.
// Callers may zero the returned slice after use.
func (g *Generator) TokenBase64Bytes(nBytes int) ([]byte, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return nil, err
	}
	out := make([]byte, base64.StdEncoding.EncodedLen(len(b)))
	base64.StdEncoding.Encode(out, b)
	return out, nil
}

// TokenURLSafe returns a URL-safe base64 string without padding
// encoding nBytes of random data.
// Note: strings are immutable; use TokenURLSafeBytes if you need to wipe secrets.
//
// Parameters:
//   - nBytes: The number of random bytes to generate.
//
// Returns:
//   - string: A URL-safe base64 string without padding.
//   - error: An error if nBytes < 0 or if entropy fails.
func (g *Generator) TokenURLSafe(nBytes int) (string, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

// TokenURLSafeBytes returns a URL-safe base64 token as a byte slice.
// Callers may zero the returned slice after use.
func (g *Generator) TokenURLSafeBytes(nBytes int) ([]byte, error) {
	b, err := g.rng.Bytes(nBytes)
	if err != nil {
		return nil, err
	}
	out := make([]byte, base64.RawURLEncoding.EncodedLen(len(b)))
	base64.RawURLEncoding.Encode(out, b)
	return out, nil
}

// StringSlice returns a slice of random strings with per-item length in
// [minStrLen, maxStrLen], using the generator's entropy source.
func (g *Generator) StringSlice(
	sliceLength, minStrLen, maxStrLen int,
) ([]string, error) {
	if sliceLength < 0 {
		return nil, core.ErrNegativeLength
	}
	result := make([]string, sliceLength)
	for i := 0; i < sliceLength; i++ {
		sLen, err := g.rng.IntRange(minStrLen, maxStrLen)
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

func isASCIICharset(charset string) bool {
	for i := 0; i < len(charset); i++ {
		if charset[i] > 0x7f {
			return false
		}
	}
	return true
}

func u64ToInt(v uint64) (int, error) {
	maxInt := uint64(^uint(0) >> 1)
	if v > maxInt {
		return 0, core.ErrResultOutOfRange
	}
	return int(v), nil
}

func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}

// Predefined character sets.
const (
	lowerCase           = "abcdefghijklmnopqrstuvwxyz"
	numbers             = "0123456789"
	lowerCaseAndNumbers = lowerCase + numbers
)
