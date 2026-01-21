package core

import (
	"math"
	"testing"
)

func TestAbsInt64ToUint64(t *testing.T) {
	tests := []struct {
		in   int64
		want uint64
	}{
		{-1, 1},
		{0, 0},
		{1, 1},
		{math.MinInt64, uint64(1) << 63},
		{math.MaxInt64, (uint64(1) << 63) - 1},
	}
	for _, tc := range tests {
		got := absInt64ToUint64(tc.in)
		if got != tc.want {
			t.Fatalf("absInt64ToUint64(%d)=%d want %d", tc.in, got, tc.want)
		}
	}
}
