package uuid

import (
	"errors"
	"regexp"
)

// UUID is a lower-case canonical textual UUID.
type UUID string

// We keep two regexes:
// - reCanon accepts canonical 8-4-4-4-12 with either hex case.
// - reLower asserts our canonical lower-case invariant.
var (
	reCanon = regexp.MustCompile(
		`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-` +
			`[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`,
	)
	reLower = regexp.MustCompile(
		`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-` +
			`[0-9a-f]{4}-[0-9a-f]{12}$`,
	)
	nilUUID = UUID("00000000-0000-0000-0000-000000000000")
)

// V4 returns a RFC 4122, variant 1 UUID v4.
// It reads 16 bytes, then sets version and variant bits.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
//   - error: An error if crypto/rand fails.
func V4() (UUID, error) {
	return Default.V4()
}

// MustV4 returns a v4 UUID or panics.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
func MustV4() UUID {
	u, err := V4()
	if err != nil {
		panic(err)
	}
	return u
}

// V7 returns a RFC 9562, variant 1 UUID v7 (time-ordered).
// The first 48 bits encode Unix milliseconds big-endian.
//
// Returns:
//   - UUID: A random UUID conforming to Version 7 and Variant 1.
//   - error: An error if crypto/rand fails.
func V7() (UUID, error) {
	return Default.V7()
}

// MustV7 returns a v7 UUID or panics.
//
// Returns:
//   - UUID: A random UUID conforming to Version 7 and Variant 1.
func MustV7() UUID {
	u, err := V7()
	if err != nil {
		panic(err)
	}
	return u
}

// Parse validates s (canonical 8-4-4-4-12, any case) and returns a
// lower-case UUID.
//
// Parameters:
//   - s: The string to parse.
//
// Returns:
//   - UUID: A lower-case UUID.
//   - error: An error if the string is invalid.
func Parse(s string) (UUID, error) {
	if !reCanon.MatchString(s) {
		return "", errors.New("invalid UUID format")
	}
	return UUID(toLowerASCII(s)), nil
}

// MustParse panics on invalid input.
//
// Parameters:
//   - s: The string to parse.
//
// Returns:
//   - UUID: A lower-case UUID.
//   - error: An error if the string is invalid.
func MustParse(s string) UUID {
	u, err := Parse(s)
	if err != nil {
		panic(err)
	}
	return u
}

// Nil returns the canonical nil UUID.
//
// Returns:
//   - UUID: The canonical nil UUID.
func Nil() UUID { return nilUUID }

// IsNil reports whether u equals the nil UUID.
//
// Parameters:
//   - u: The UUID to check.
//
// Returns:
//   - bool: True if u equals the nil UUID.
func (u UUID) IsNil() bool { return u == nilUUID }

// String returns the textual UUID.
//
// Parameters:
//   - u: The UUID to return the string representation of.
//
// Returns:
//   - string: The string representation of the UUID.
func (u UUID) String() string { return string(u) }

// Bytes returns the 16-byte representation of a canonical lower-case
// UUID. It rejects non-lower-case or malformed values.
//
// Parameters:
//   - u: The UUID to return the 16-byte representation of.
//
// Returns:
//   - [16]byte: The 16-byte representation of the UUID.
//   - error: An error if the UUID is invalid.
func (u UUID) Bytes() ([16]byte, error) {
	var out [16]byte
	s := string(u)
	if !reLower.MatchString(s) {
		return out, errors.New("invalid UUID")
	}
	// Decode two hex nibbles at a time, skipping hyphens.
	di := 0 // dest index in out
	for i := 0; i < len(s); {
		if s[i] == '-' {
			i++
			continue
		}
		if i+1 >= len(s) || di >= 16 {
			return out, errors.New("invalid UUID length")
		}
		hi := fromHexNibble(s[i])
		lo := fromHexNibble(s[i+1])
		if hi == 255 || lo == 255 {
			return out, errors.New("invalid UUID hex")
		}
		out[di] = (hi << 4) | lo
		di++
		i += 2
	}
	if di != 16 {
		return out, errors.New("invalid UUID length")
	}
	return out, nil
}

// fromBytes formats a 16-byte slice into canonical lower-case string.
func fromBytes(b []byte) UUID {
	dst := make([]byte, 36)
	// Hyphen positions.
	hy := map[int]struct{}{8: {}, 13: {}, 18: {}, 23: {}}
	si, di := 0, 0
	for di < 36 {
		if _, ok := hy[di]; ok {
			dst[di] = '-'
			di++
			continue
		}
		dst[di], dst[di+1] = hexHiLo(b[si])
		si++
		di += 2
	}
	return UUID(string(dst))
}

func hexHiLo(v byte) (byte, byte) {
	const hex = "0123456789abcdef"
	return hex[v>>4], hex[v&0x0f]
}

func toLowerASCII(s string) string {
	b := []byte(s)
	for i := range b {
		if 'A' <= b[i] && b[i] <= 'Z' {
			b[i] = b[i] + 'a' - 'A'
		}
	}
	return string(b)
}

func fromHexNibble(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	default:
		return 255
	}
}
