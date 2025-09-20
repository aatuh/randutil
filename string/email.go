package string

import (
	"errors"
	"fmt"
)

// Email returns a random email of exactly totalLength chars in the form
// local@domain.com (5 chars reserved for "@" + ".com"). totalLength
// must be at least 7 so both local and domain have >= 1 char.
//
// Parameters:
//   - totalLength: The total length of the email address (must be at least 7).
//
// Returns:
//   - string: A random email address.
//   - error: An error if crypto/rand fails.
func Email(totalLength int) (string, error) {
	if totalLength < 7 {
		return "", errors.New("totalLength must be at least 7")
	}
	// Ensure exact total length and at least one char per side.
	body := totalLength - 5
	localLen := body / 2
	if localLen == 0 {
		localLen = 1
	}
	domainLen := body - localLen
	if domainLen == 0 {
		domainLen = 1
		localLen = body - domainLen
	}
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

// MustEmail returns a random email address of totalLength characters. It panics
// on error.
//
// Parameters:
//   - totalLength: The total length of the email address (must be at least 7).
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
