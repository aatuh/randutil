package core

import "errors"

// Package-level errors. Returned when arguments are invalid.
var (
	ErrInvalidN        = errors.New("randutil: n must be > 0")
	ErrInvalidRange    = errors.New("randutil: range must be > 0")
	ErrSampleTooLarge  = errors.New("randutil: sample k exceeds size")
	ErrInvalidWeights  = errors.New("randutil: invalid weights")
	ErrWeightsMismatch = errors.New("randutil: items/weights mismatch")
)
