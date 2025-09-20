package numeric

import (
	"encoding/binary"

	"github.com/aatuh/randutil/core"
)

// Uint64 returns a random uint64 from the active source.
//
// Returns:
//   - uint64: A random uint64 value.
//   - error: An error if the source fails.
func Uint64() (uint64, error) {
	var b [8]byte
	if err := readFull(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

// Uint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - uint64: A random uint64 in [0, n).
//   - error: An error if the source fails or n == 0.
func Uint64n(n uint64) (uint64, error) {
	if n == 0 {
		return 0, core.ErrInvalidRange
	}
	var (
		max   = ^uint64(0)
		limit = max - (max % n)
	)
	for {
		u, err := Uint64()
		if err != nil {
			return 0, err
		}
		if u < limit {
			return u % n, nil
		}
	}
}

// Intn returns a uniform random int in [0, n). n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - int: A random int in [0, n).
//   - error: An error if the source fails or n <= 0.
func Intn(n int) (int, error) {
	if n <= 0 {
		return 0, core.ErrInvalidN
	}
	u, err := Uint64n(uint64(n))
	if err != nil {
		return 0, err
	}
	return int(u), nil
}

// Int64n returns a uniform random int64 in [0, n). n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - int64: A random int64 in [0, n).
//   - error: An error if the source fails or n <= 0.
func Int64n(n int64) (int64, error) {
	if n <= 0 {
		return 0, core.ErrInvalidN
	}
	u, err := Uint64n(uint64(n))
	if err != nil {
		return 0, err
	}
	return int64(u), nil
}

// Float64 returns a uniform random float64 in [0.0, 1.0) with 53 bits
// of precision built from the active entropy source.
//
// Returns:
//   - float64: A random float64 in [0.0, 1.0).
//   - error: An error if the source fails.
func Float64() (float64, error) {
	var b [8]byte
	if err := readFull(b[:]); err != nil {
		return 0, err
	}
	u := binary.LittleEndian.Uint64(b[:]) >> 11
	const denom = 1 << 53
	return float64(u) / float64(denom), nil
}

// MustUint64 returns a random uint64 from crypto/rand. It panics on error.
//
// Returns:
//   - uint64: A random uint64 value.
func MustUint64() uint64 {
	u, err := Uint64()
	if err != nil {
		panic(err)
	}
	return u
}

// MustUint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - uint64: A random uint64 in [0, n).
func MustUint64n(n uint64) uint64 {
	u, err := Uint64n(n)
	if err != nil {
		panic(err)
	}
	return u
}

// MustIntn returns a uniform random int in [0, n). It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - int: A random int in [0, n).
func MustIntn(n int) int {
	v, err := Intn(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustInt64n returns a uniform random int64 in [0, n). It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - int64: A random int64 in [0, n).
func MustInt64n(n int64) int64 {
	v, err := Int64n(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustFloat64 returns a uniform random float64 in [0.0, 1.0) with 53 bits
// of precision. It panics on error.
//
// Returns:
//   - float64: A random float64 in [0.0, 1.0).
func MustFloat64() float64 {
	v, err := Float64()
	if err != nil {
		panic(err)
	}
	return v
}
