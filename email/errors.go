package email

import "errors"

// Package-level errors for email validation.
var (
	ErrTotalLengthTooSmall = errors.New("randutil: total length must be >= 7")
)
