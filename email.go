package randutil

import (
	"errors"
	"fmt"
)

// Email returns a random email address of totalLength characters. The
// email is built as local@domain.com (5 characters reserved for "@" and
// ".com"). totalLength must be at least 6.
//
// Returns:
//   - string: A random email address.
//   - error: An error if crypto/rand fails.
func Email(totalLength int) (string, error) {
	if totalLength < 6 {
		return "", errors.New("totalLength must be at least 6")
	}
	localLen := (totalLength - 5) / 2
	domainLen := totalLength - 5 - localLen
	local, err := String(localLen)
	if err != nil {
		return "", err
	}
	domain, err := String(domainLen)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s.com", local, domain), nil
}

// Email returns a random email address of totalLength characters. It panics
// on error.
//
// Returns:
//   - string: A random email address.
func MustEmail(totalLength int) string {
	s, err := Email(totalLength)
	if err != nil {
		panic(err)
	}
	return s
}
