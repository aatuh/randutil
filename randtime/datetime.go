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
	return Default().Datetime()
}

// TimeInNearPast returns a time a few minutes in the past.
//
// Returns:
//   - time.Time: A time a few minutes in the past.
//   - error: An error if crypto/rand fails.
func TimeInNearPast() (time.Time, error) {
	return Default().TimeInNearPast()
}

// TimeInNearFuture returns a time a few minutes in the future.
//
// Returns:
//   - time.Time: A time a few minutes in the future.
//   - error: An error if crypto/rand fails.
func TimeInNearFuture() (time.Time, error) {
	return Default().TimeInNearFuture()
}
