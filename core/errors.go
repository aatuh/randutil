package core

import "errors"

// Package-level errors. Returned when arguments are invalid.
var (
	ErrInvalidN        = errors.New("randutil: n must be > 0")
	ErrInvalidRange    = errors.New("randutil: range must be > 0")
	ErrSampleTooLarge  = errors.New("randutil: sample k exceeds size")
	ErrInvalidWeights  = errors.New("randutil: invalid weights")
	ErrWeightsMismatch = errors.New("randutil: items/weights mismatch")
	ErrEmptySlice      = errors.New("randutil: cannot pick from empty slice")
	ErrEmptyItems      = errors.New("randutil: empty items")
	ErrEmptyCharset    = errors.New("randutil: charset must be non-empty")
	ErrOddHexLength    = errors.New("randutil: hex length must be even")

	// Range and generation errors
	ErrMinGreaterThanMax       = errors.New("randutil: min > max")
	ErrInvalidRangeNonPositive = errors.New("randutil: non-positive range")
	ErrResultOutOfRange        = errors.New("randutil: result out of range")
)
