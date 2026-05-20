package ulid

import (
	"errors"
	"strings"
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
	upper := strings.ToUpper(s)
	if _, err := ulidEncoding.DecodeString(upper); err != nil {
		return "", ErrInvalidULID
	}
	return ULID(upper), nil
}

// String returns the ULID string.
func (u ULID) String() string { return string(u) }
