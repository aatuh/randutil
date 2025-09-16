package randutil

import (
	"time"
)

// Datetime returns a secure random time between year 1 and 9999.
//
// Returns:
//   - time.Time: A random time between year 1 and 9999.
//   - error: An error if crypto/rand fails.
func Datetime() (time.Time, error) {
	year, err := IntRange(1, 9999)
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := IntRange(1, 12)
	if err != nil {
		return time.Time{}, err
	}
	month := time.Month(monthInt)
	day, err := IntRange(1, daysInMonth(year, month))
	if err != nil {
		return time.Time{}, err
	}
	hour, err := IntRange(0, 23)
	if err != nil {
		return time.Time{}, err
	}
	minute, err := IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	second, err := IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	nano, err := IntRange(0, int(time.Second)-1)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, month, day, hour, minute, second, nano,
		time.UTC), nil
}

// Datetime returns a secure random time between year 1 and 9999. It panics
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
	offset, err := IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().UTC().Add(-time.Minute *
		time.Duration(offset)), nil
}

// TimeInNearPast returns a time a few minutes in the past. It panics if an
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
	offset, err := IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().UTC().Add(time.Minute *
		time.Duration(offset)), nil
}

// TimeInNearFuture returns a time a few minutes in the future. It panics if
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

// daysInMonth returns the number of days in the given month of year.
func daysInMonth(year int, month time.Month) int {
	if month == 2 {
		if isLeapYear(year) {
			return 29
		}
		return 28
	}
	switch month {
	case 4, 6, 9, 11:
		return 30
	default:
		return 31
	}
}

// isLeapYear returns true if year is a leap year.
func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}
