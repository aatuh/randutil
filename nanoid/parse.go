package nanoid

import (
	"errors"

	"github.com/aatuh/randutil/v2/core"
)

// ErrInvalidID is returned when parsing fails.
var ErrInvalidID = errors.New("randutil: invalid nanoid")

// Parse validates id against the default alphabet.
func Parse(id string) (string, error) {
	return ParseWithAlphabet(id, DefaultAlphabet)
}

// ParseWithAlphabet validates id against alphabet.
func ParseWithAlphabet(id string, alphabet string) (string, error) {
	if id == "" {
		return "", ErrInvalidID
	}
	if alphabet == "" {
		return "", core.ErrEmptyCharset
	}
	if !isASCIIAlphabet(alphabet) {
		return "", core.ErrInvalidCharset
	}
	var allowed [128]bool
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] < 0x80 {
			allowed[alphabet[i]] = true
		}
	}
	for i := 0; i < len(id); i++ {
		b := id[i]
		if b > 0x7f || !allowed[b] {
			return "", ErrInvalidID
		}
	}
	return id, nil
}

func isASCIIAlphabet(alphabet string) bool {
	for i := 0; i < len(alphabet); i++ {
		if alphabet[i] > 0x7f {
			return false
		}
	}
	return true
}
