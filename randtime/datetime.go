package randtime

import (
	"time"
)

// Datetime returns a secure random time between year 1 and 9999.
//
// Returns:
//   - time.Time: A random time between year 1 and 9999.
//   - error: An error if crypto/rand fails.
func Datetime() (time.Time, error) {
	return Default.Datetime()
}

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

// TimeInNearPast returns a time a few minutes in the past.
//
// Returns:
//   - time.Time: A time a few minutes in the past.
//   - error: An error if crypto/rand fails.
func TimeInNearPast() (time.Time, error) {
	return Default.TimeInNearPast()
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

// TimeInNearFuture returns a time a few minutes in the future.
//
// Returns:
//   - time.Time: A time a few minutes in the future.
//   - error: An error if crypto/rand fails.
func TimeInNearFuture() (time.Time, error) {
	return Default.TimeInNearFuture()
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
