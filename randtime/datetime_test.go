package randtime

import (
	"testing"
	stdtime "time"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestDatetimeBounds(t *testing.T) {
	for i := 0; i < 50; i++ {
		v, err := Datetime()
		if err != nil {
			t.Fatalf("Datetime error: %v", err)
		}
		year := v.Year()
		if year < 1 || year > 9999 {
			t.Fatalf("Datetime year %d out of range", year)
		}
		if v.Month() < 1 || v.Month() > 12 {
			t.Fatalf("Datetime month %d out of range", v.Month())
		}
		dm := daysInMonth(year, v.Month())
		if v.Day() < 1 || v.Day() > dm {
			t.Fatalf("Datetime day %d out of range (month has %d)", v.Day(), dm)
		}
	}
}

func TestTimeInNearPastFuture(t *testing.T) {
	past, err := TimeInNearPast()
	if err != nil {
		t.Fatalf("TimeInNearPast error: %v", err)
	}
	future, err := TimeInNearFuture()
	if err != nil {
		t.Fatalf("TimeInNearFuture error: %v", err)
	}
	now := stdtime.Now().UTC()
	if !(past.Before(now) || past.Equal(now)) {
		t.Fatalf("TimeInNearPast returned future time: %v >= %v", past, now)
	}
	if !(future.After(now) || future.Equal(now)) {
		t.Fatalf("TimeInNearFuture returned past time: %v <= %v", future, now)
	}
	const (
		minDelta = 5*stdtime.Minute - 2*stdtime.Second
		maxDelta = 10*stdtime.Minute + 2*stdtime.Second
	)
	if delta := now.Sub(past); delta < minDelta || delta > maxDelta {
		t.Fatalf("TimeInNearPast delta %v outside expected range", delta)
	}
	if delta := future.Sub(now); delta < minDelta || delta > maxDelta {
		t.Fatalf("TimeInNearFuture delta %v outside expected range", delta)
	}
}

func TestTimeInNearPastFutureWithClock(t *testing.T) {
	fixed := stdtime.Date(2024, 1, 2, 3, 4, 5, 0, stdtime.UTC)
	src := testutil.NewSeqReader([]byte{0})
	gen := NewWithClock(core.New(src), func() stdtime.Time { return fixed })

	past, err := gen.TimeInNearPast()
	if err != nil {
		t.Fatalf("TimeInNearPast error: %v", err)
	}
	wantPast := fixed.Add(-5 * stdtime.Minute)
	if !past.Equal(wantPast) {
		t.Fatalf("TimeInNearPast=%v want %v", past, wantPast)
	}

	future, err := gen.TimeInNearFuture()
	if err != nil {
		t.Fatalf("TimeInNearFuture error: %v", err)
	}
	wantFuture := fixed.Add(5 * stdtime.Minute)
	if !future.Equal(wantFuture) {
		t.Fatalf("TimeInNearFuture=%v want %v", future, wantFuture)
	}
}

func TestDaysInMonth(t *testing.T) {
	tests := []struct {
		year  int
		month stdtime.Month
		want  int
	}{
		{2023, stdtime.January, 31},
		{2024, stdtime.February, 29},
		{1900, stdtime.February, 28},
		{2000, stdtime.February, 29},
		{2023, stdtime.April, 30},
	}
	for _, tc := range tests {
		if got := daysInMonth(tc.year, tc.month); got != tc.want {
			t.Fatalf("daysInMonth(%d,%d)=%d want %d", tc.year, tc.month, got, tc.want)
		}
	}
}

func TestIsLeapYear(t *testing.T) {
	cases := map[int]bool{
		1600: true,
		1700: false,
		2000: true,
		2020: true,
		2023: false,
	}
	for year, want := range cases {
		if got := isLeapYear(year); got != want {
			t.Fatalf("isLeapYear(%d)=%v want %v", year, got, want)
		}
	}
}
