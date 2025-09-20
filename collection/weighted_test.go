package collection

import (
	"math"
	"sort"
	"testing"
)

func TestWeightedChoiceAndSample(t *testing.T) {
	items := []string{"a", "b", "c"}
	w := []float64{0, 1, 2}

	// Choice should never return zero-weight "a".
	for i := 0; i < 100; i++ {
		x, err := WeightedChoice(items, w)
		if err != nil {
			t.Fatalf("WeightedChoice error: %v", err)
		}
		if x == "a" {
			t.Fatalf("picked zero-weight item")
		}
	}

	// Sample k=2 should contain only from {"b","c"} and be unique.
	s, err := WeightedSample(items, w, 2)
	if err != nil {
		t.Fatalf("WeightedSample error: %v", err)
	}
	if len(s) != 2 {
		t.Fatalf("len(sample)=%d want 2", len(s))
	}
	if s[0] == s[1] {
		t.Fatalf("duplicate items in sample")
	}
	for _, v := range s {
		if v == "a" {
			t.Fatalf("sample contains zero-weight item")
		}
	}
}

func TestWeightedChoiceErrors(t *testing.T) {
	items := []string{"a", "b"}
	if _, err := WeightedChoice(items, []float64{}); err != ErrWeightsMismatch {
		t.Fatalf("expected weights mismatch error, got %v", err)
	}
	if _, err := WeightedChoice([]string{}, []float64{}); err == nil {
		t.Fatalf("expected error for empty items")
	}
	if _, err := WeightedChoice(items, []float64{-1, 2}); err != ErrInvalidWeights {
		t.Fatalf("expected invalid weights, got %v", err)
	}
	if _, err := WeightedChoice(items, []float64{0, 0}); err != ErrInvalidWeights {
		t.Fatalf("expected invalid weights for zero sum, got %v", err)
	}
}

func TestWeightedChoiceRespectsZeroWeights(t *testing.T) {
	items := []int{1, 2}
	weights := []float64{0, 1}
	for i := 0; i < 64; i++ {
		v, err := WeightedChoice(items, weights)
		if err != nil {
			t.Fatalf("WeightedChoice error: %v", err)
		}
		if v == 1 {
			t.Fatalf("WeightedChoice selected zero-weight item: %d", v)
		}
	}
}

func TestWeightedSampleErrors(t *testing.T) {
	items := []int{1, 2}
	weights := []float64{0.5, 0.5}
	if _, err := WeightedSample(items, weights, -1); err == nil {
		t.Fatalf("expected error for negative k")
	}
	if _, err := WeightedSample([]int{}, []float64{}, 1); err == nil {
		t.Fatalf("expected error for empty items")
	}
	if _, err := WeightedSample(items, []float64{0.5}, 1); err != ErrWeightsMismatch {
		t.Fatalf("expected weights mismatch, got %v", err)
	}
	if _, err := WeightedSample(items, []float64{-1, 1}, 1); err != ErrInvalidWeights {
		t.Fatalf("expected invalid weights, got %v", err)
	}
	if _, err := WeightedSample(items, []float64{0, 0}, 1); err == nil || err.Error() != "sample k exceeds size" {
		t.Fatalf("expected sample too large error, got %v", err)
	}
}

func TestWeightedSampleReturnsDistinct(t *testing.T) {
	items := []string{"a", "b", "c", "d"}
	weights := []float64{1, 2, 3, 4}
	out, err := WeightedSample(items, weights, 3)
	if err != nil {
		t.Fatalf("WeightedSample error: %v", err)
	}
	if len(out) != 3 {
		t.Fatalf("WeightedSample len = %d want 3", len(out))
	}
	seen := map[string]bool{}
	for _, v := range out {
		if seen[v] {
			t.Fatalf("WeightedSample duplicated item %q", v)
		}
		seen[v] = true
	}
	for _, v := range out {
		if indexOf(items, v) == -1 {
			t.Fatalf("WeightedSample produced unknown item %q", v)
		}
	}
}

func TestFloat64ReturnsUnitInterval(t *testing.T) {
	for i := 0; i < 256; i++ {
		v, err := Float64()
		if err != nil {
			t.Fatalf("Float64 error: %v", err)
		}
		if v < 0 || v >= 1 {
			t.Fatalf("Float64 value %f out of range", v)
		}
	}
}

// indexOf returns the index of v in s or -1.
func indexOf[T comparable](s []T, v T) int {
	for i, item := range s {
		if item == v {
			return i
		}
	}
	return -1
}

func TestWeightedSampleFavoursLargerWeights(t *testing.T) {
	items := []int{1, 2, 3}
	weights := []float64{0, 5, 0}
	out, err := WeightedSample(items, weights, 1)
	if err != nil {
		t.Fatalf("WeightedSample error: %v", err)
	}
	if len(out) != 1 || out[0] != 2 {
		t.Fatalf("WeightedSample expected only index 2, got %v", out)
	}
}

func TestWeightedSampleStableOrdering(t *testing.T) {
	items := []int{1, 2, 3, 4}
	weights := []float64{1, 1, 1, 1}
	out, err := WeightedSample(items, weights, 4)
	if err != nil {
		t.Fatalf("WeightedSample error: %v", err)
	}
	sorted := append([]int(nil), out...)
	sort.Ints(sorted)
	for i, v := range sorted {
		if v != i+1 {
			t.Fatalf("WeightedSample missing value %d in %v", i+1, sorted)
		}
	}
}

func TestFloat64HighPrecision(t *testing.T) {
	// Ensure we exercise more than one code path and values look random.
	sum := 0.0
	for i := 0; i < 1024; i++ {
		v, err := Float64()
		if err != nil {
			t.Fatalf("Float64 error: %v", err)
		}
		sum += v
	}
	mean := sum / 1024
	if math.Abs(mean-0.5) > 0.2 {
		t.Fatalf("Float64 mean too far from 0.5: %f", mean)
	}
}
