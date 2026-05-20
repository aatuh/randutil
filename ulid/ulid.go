package ulid

import (
	"errors"
)

// ULID is the canonical string form of a ULID.
type ULID string

// ErrInvalidULID is returned when parsing fails.
var ErrInvalidULID = errors.New("randutil: invalid ulid")

// ID returns a new ULID from the default generator.
func ID() (ULID, error) {
	return Default().ULID()
}

// Parse validates s and returns a canonical upper-case ULID.
func Parse(s string) (ULID, error) {
	if len(s) != encodedLen {
		return "", ErrInvalidULID
	}
	var b [encodedLen]byte
	for i := 0; i < encodedLen; i++ {
		c := s[i]
		if c >= 'a' && c <= 'z' {
			c -= 'a' - 'A'
		}
		if !isULIDChar(c) {
			return "", ErrInvalidULID
		}
		b[i] = c
	}
	if b[0] > '7' {
		return "", ErrInvalidULID
	}
	upper := string(b[:])
	if _, err := ulidEncoding.DecodeString(upper); err != nil {
		return "", ErrInvalidULID
	}
	return ULID(upper), nil
}

// String returns the ULID string.
func (u ULID) String() string { return string(u) }

func isULIDChar(c byte) bool {
	switch {
	case c >= '0' && c <= '9':
		return true
	case c >= 'A' && c <= 'H':
		return true
	case c >= 'J' && c <= 'K':
		return true
	case c >= 'M' && c <= 'N':
		return true
	case c >= 'P' && c <= 'T':
		return true
	case c >= 'V' && c <= 'Z':
		return true
	default:
		return false
	}
}
