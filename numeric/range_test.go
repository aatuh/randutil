package numeric

import (
	"math"
	"strconv"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
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
		t.Skip("AnyInt test skipped on 64-bit platforms to avoid large range issues")
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

func TestGeneratorAnyInt64CanReturnMinInt64(t *testing.T) {
	gen := New(core.New(testutil.NewSeqReader(testutil.Uint64Bytes(0))))
	got, err := gen.AnyInt64()
	if err != nil {
		t.Fatalf("AnyInt64 returned error: %v", err)
	}
	if got != minInt64 {
		t.Fatalf("AnyInt64 = %d want %d", got, minInt64)
	}
}

func TestGeneratorAnyIntCanReturnMinIntOn64Bit(t *testing.T) {
	if strconv.IntSize != 64 {
		t.Skip("64-bit AnyInt regression")
	}
	gen := New(core.New(testutil.NewSeqReader(testutil.Uint64Bytes(0))))
	got, err := gen.AnyInt()
	if err != nil {
		t.Fatalf("AnyInt returned error: %v", err)
	}
	if got != minInt {
		t.Fatalf("AnyInt = %d want %d", got, minInt)
	}
}
