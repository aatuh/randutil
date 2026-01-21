package core

import "errors"

// Package-level errors. Returned when arguments are invalid.
var (
	ErrNegativeLength     = errors.New("randutil: length must be >= 0")
	ErrNonPositiveBound   = errors.New("randutil: bound must be > 0")
	ErrInvalidProbability = errors.New("randutil: probability must be in [0,1]")
	ErrNonPositiveRate    = errors.New("randutil: rate must be > 0")
	ErrNegativeStdDev     = errors.New("randutil: stddev must be >= 0")
	ErrEmptyCharset       = errors.New("randutil: charset must be non-empty")
	ErrInvalidCharset     = errors.New("randutil: charset must be ASCII")
	ErrOddHexLength       = errors.New("randutil: hex length must be even")

	ErrSampleTooLarge  = errors.New("randutil: sample size exceeds available items")
	ErrInvalidWeights  = errors.New("randutil: weights must be non-negative with at least one > 0")
	ErrWeightsMismatch = errors.New("randutil: items/weights length mismatch")
	ErrEmptySlice      = errors.New("randutil: empty slice")
	ErrEmptyItems      = errors.New("randutil: empty items")

	ErrMinGreaterThanMax       = errors.New("randutil: min greater than max")
	ErrInvalidRangeNonPositive = errors.New("randutil: range must be positive")
	ErrResultOutOfRange        = errors.New("randutil: result out of range")
)
