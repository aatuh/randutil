package numeric

import (
	"math"
	"strconv"
	"testing"
)

func TestIntRangeBounds(t *testing.T) {
	for i := 0; i < 100; i++ {
		v, err := IntRange(10, 20)
		if err != nil {
			t.Fatalf("IntRange error: %v", err)
		}
		if v < 10 || v > 20 {
			t.Fatalf("IntRange value %d out of bounds", v)
		}
	}
}

func TestIntRangeInvalid(t *testing.T) {
	if _, err := IntRange(5, 4); err == nil {
		t.Fatal("expected error when min > max")
	}
}

func TestAnyIntInBounds(t *testing.T) {
	if strconv.IntSize > 32 {
		t.Skip("AnyInt uses 32-bit math; skip on 64-bit platforms")
	}
	v, err := AnyInt()
	if err != nil {
		t.Fatalf("AnyInt returned error: %v", err)
	}
	if v < math.MinInt || v > math.MaxInt {
		t.Fatalf("AnyInt returned %d outside int range", v)
	}
}

func TestInt32RangeAndVariants(t *testing.T) {
	if v, err := Int32Range(-2, 2); err != nil || v < -2 || v > 2 {
		t.Fatalf("Int32Range value: %d err: %v", v, err)
	}
	if v, err := PositiveInt32(); err != nil || v < 1 {
		t.Fatalf("PositiveInt32 value: %d err: %v", v, err)
	}
	if v, err := NegativeInt32(); err != nil || v > -1 {
		t.Fatalf("NegativeInt32 value: %d err: %v", v, err)
	}
}

func TestInt64RangeAndVariants(t *testing.T) {
	if v, err := Int64Range(-5, 5); err != nil || v < -5 || v > 5 {
		t.Fatalf("Int64Range value: %d err: %v", v, err)
	}
	if v, err := PositiveInt64(); err != nil || v < 1 {
		t.Fatalf("PositiveInt64 value: %d err: %v", v, err)
	}
	if v, err := NegativeInt64(); err != nil || v > -1 {
		t.Fatalf("NegativeInt64 value: %d err: %v", v, err)
	}
}
