//go:build randutil_must
// +build randutil_must

package randtime

import "time"

// MustDatetime returns a secure random time between year 1 and 9999. It panics
// if an error occurs.
//
// Returns:
//   - time.Time: A random time between year 1 and 9999.
func MustDatetime() time.Time {
	t, err := Datetime()
	if err != nil {
		panic(err)
	}
	return t
}

// MustTimeInNearPast returns a time a few minutes in the past. It panics if an
// error occurs.
//
// Returns:
//   - time.Time: A time a few minutes in the past.
func MustTimeInNearPast() time.Time {
	t, err := TimeInNearPast()
	if err != nil {
		panic(err)
	}
	return t
}

// MustTimeInNearFuture returns a time a few minutes in the future. It panics if
// an error occurs.
//
// Returns:
//   - time.Time: A time a few minutes in the future.
func MustTimeInNearFuture() time.Time {
	t, err := TimeInNearFuture()
	if err != nil {
		panic(err)
	}
	return t
}

// MustJitter returns base adjusted by a random factor in [-pct, +pct].
func MustJitter(base time.Duration, pct float64) time.Duration {
	d, err := Jitter(base, pct)
	if err != nil {
		panic(err)
	}
	return d
}
