package randtime

import (
	"time"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds random time values using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng core.RNG
	now func() time.Time
}

// New returns a time Generator. If rng is nil, crypto/rand is used.
func New(rng core.RNG) *Generator {
	return NewWithClock(rng, time.Now)
}

// NewWithClock returns a time Generator bound to rng and clock.
// If rng is nil, crypto/rand is used. If now is nil, time.Now is used.
func NewWithClock(rng core.RNG, now func() time.Time) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	if now == nil {
		now = time.Now
	}
	return &Generator{rng: rng, now: now}
}

// NewWithSource returns a time Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return NewWithClock(core.New(src), time.Now)
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// Datetime returns a secure random time between year 1 and 9999.
//
// Returns:
//   - time.Time: A random time between year 1 and 9999.
//   - error: An error if entropy fails.
func (g *Generator) Datetime() (time.Time, error) {
	year, err := g.rng.IntRange(1, 9999)
	if err != nil {
		return time.Time{}, err
	}
	monthInt, err := g.rng.IntRange(1, 12)
	if err != nil {
		return time.Time{}, err
	}
	month := time.Month(monthInt)
	day, err := g.rng.IntRange(1, daysInMonth(year, month))
	if err != nil {
		return time.Time{}, err
	}
	hour, err := g.rng.IntRange(0, 23)
	if err != nil {
		return time.Time{}, err
	}
	minute, err := g.rng.IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	second, err := g.rng.IntRange(0, 59)
	if err != nil {
		return time.Time{}, err
	}
	nano, err := g.rng.IntRange(0, int(time.Second)-1)
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
	offset, err := g.rng.IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return g.nowUTC().Add(-time.Minute * time.Duration(offset)), nil
}

// TimeInNearFuture returns a time a few minutes in the future.
//
// Returns:
//   - time.Time: A random time 5-10 minutes in the future.
//   - error: An error if entropy fails.
func (g *Generator) TimeInNearFuture() (time.Time, error) {
	offset, err := g.rng.IntRange(5, 10)
	if err != nil {
		return time.Time{}, err
	}
	return g.nowUTC().Add(time.Minute * time.Duration(offset)), nil
}

func (g *Generator) nowUTC() time.Time {
	if g == nil || g.now == nil {
		return time.Now().UTC()
	}
	return g.now().UTC()
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
