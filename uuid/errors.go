package uuid

import "errors"

// Package-level errors for UUID validation.
var (
	ErrInvalidFormat = errors.New("randutil: invalid UUID format")
	ErrInvalidUUID   = errors.New("randutil: invalid UUID")
)
