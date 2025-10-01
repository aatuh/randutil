package randtime

import (
	"io"
	"time"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds random time values using a core generator.
type Generator struct {
	G core.Generator
}

// New returns a time Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator { return &Generator{G: core.Generator{R: src}} }

// Default is the package-wide default generator.
var Default = New(nil)

// Datetime returns a secure random time between year 1 and 9999.
//
// Returns:
//   - time.Time: A random time between year 1 and 9999.
//   - error: An error if entropy fails.
func (g *Generator) Datetime() (time.Time, error) {
	year, err := g.G.IntRange(1, 9999)
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := g.G.IntRange(1, 12)
	if err != nil {
		return time.Time{}, err
	}
	month := time.Month(monthInt)
	day, err := g.G.IntRange(1, daysInMonth(year, month))
	if err != nil {
		return time.Time{}, err
	}
	hour, err := g.G.IntRange(0, 23)
	if err != nil {
		return time.Time{}, err
	}
	minute, err := g.G.IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	second, err := g.G.IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	nano, err := g.G.IntRange(0, int(time.Second)-1)
	if err != nil {
		return time.Time{}, err
	}
	return time.Date(year, month, day, hour, minute, second, nano,
		time.UTC), nil
}

// TimeInNearPast returns a time a few minutes in the past.
//
// Returns:
//   - time.Time: A random time 5-10 minutes in the past.
//   - error: An error if entropy fails.
func (g *Generator) TimeInNearPast() (time.Time, error) {
	offset, err := g.G.IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().UTC().Add(-time.Minute *
		time.Duration(offset)), nil
}

// TimeInNearFuture returns a time a few minutes in the future.
//
// Returns:
//   - time.Time: A random time 5-10 minutes in the future.
//   - error: An error if entropy fails.
func (g *Generator) TimeInNearFuture() (time.Time, error) {
	offset, err := g.G.IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return time.Now().UTC().Add(time.Minute *
		time.Duration(offset)), nil
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
