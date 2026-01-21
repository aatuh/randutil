package uuid

// UUID is a lower-case canonical textual UUID.
type UUID string

const canonicalLen = 36

var nilUUID = UUID("00000000-0000-0000-0000-000000000000")

// V4 returns a RFC 4122, variant 1 UUID v4.
// It reads 16 bytes, then sets version and variant bits.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
//   - error: An error if crypto/rand fails.
func V4() (UUID, error) {
	return Default().V4()
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
	return Default().V7()
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
	if !isCanonicalUUID(s, true) {
		return "", ErrInvalidFormat
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
	if !isCanonicalUUID(s, false) {
		return out, ErrInvalidUUID
	}
	// Decode two hex nibbles at a time, skipping hyphens.
	si := 0 // source index
	for di := range out {
		for si < len(s) && s[si] == '-' {
			si++
		}
		hi := fromHexNibble(s[si])
		lo := fromHexNibble(s[si+1])
		if hi == 255 || lo == 255 {
			return out, ErrInvalidUUID
		}
		// #nosec G602 -- di is bounded by range over out.
		out[di] = (hi << 4) | lo
		si += 2
	}
	return out, nil
}

// fromBytes formats a 16-byte slice into canonical lower-case string.
func fromBytes(b []byte) UUID {
	var dst [36]byte
	di := 0
	for i := 0; i < 16; i++ {
		if di == 8 || di == 13 || di == 18 || di == 23 {
			dst[di] = '-'
			di++
		}
		hi, lo := hexHiLo(b[i])
		dst[di] = hi
		dst[di+1] = lo
		di += 2
	}
	return UUID(string(dst[:]))
}

func hexHiLo(v byte) (byte, byte) {
	const hex = "0123456789abcdef"
	return hex[v>>4], hex[v&0x0f]
}

func isCanonicalUUID(s string, allowUpper bool) bool {
	if len(s) != canonicalLen {
		return false
	}
	for i := 0; i < len(s); i++ {
		switch i {
		case 8, 13, 18, 23:
			if s[i] != '-' {
				return false
			}
		default:
			c := s[i]
			if '0' <= c && c <= '9' {
				continue
			}
			if 'a' <= c && c <= 'f' {
				continue
			}
			if allowUpper && 'A' <= c && c <= 'F' {
				continue
			}
			return false
		}
	}
	return true
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
