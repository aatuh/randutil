//go:build randutil_ci
// +build randutil_ci

package core_test

import (
	"math"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestUint64nUniform(t *testing.T) {
	gen := derivedGenerator(t, "uint64n")
	assertUniform(t, 9, 100000, func() (int, error) {
		v, err := gen.Uint64n(9)
		return int(v), err
	})
}

func TestIntRangeUniform(t *testing.T) {
	gen := derivedGenerator(t, "intrange")
	min := -3
	max := 6
	width := max - min + 1
	assertUniform(t, width, 100000, func() (int, error) {
		v, err := gen.IntRange(min, max)
		return v - min, err
	})
}

func derivedGenerator(t *testing.T, label string) *core.Generator {
	t.Helper()
	src, err := adapters.DeriveSource([]byte("uniform"), label)
	if err != nil {
		t.Fatalf("DeriveSource error: %v", err)
	}
	return core.New(src)
}

func assertUniform(t *testing.T, buckets int, samples int, draw func() (int, error)) {
	t.Helper()
	counts := make([]int, buckets)
	for i := 0; i < samples; i++ {
		v, err := draw()
		if err != nil {
			t.Fatalf("draw error: %v", err)
		}
		if v < 0 || v >= buckets {
			t.Fatalf("value %d outside [0,%d)", v, buckets)
		}
		counts[v]++
	}
	expected := float64(samples) / float64(buckets)
	tol := 6 * math.Sqrt(expected)
	for i, count := range counts {
		if math.Abs(float64(count)-expected) > tol {
			t.Fatalf("bucket %d count=%d want %0.1f ± %0.1f", i, count, expected, tol)
		}
	}
}
