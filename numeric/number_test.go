package numeric

import (
	"io"
	"math"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestUint64nRange(t *testing.T) {
	const n = 10
	for i := 0; i < 1000; i++ {
		v, err := Intn(n)
		if err != nil {
			t.Fatalf("Intn error: %v", err)
		}
		if v < 0 || v >= n {
			t.Fatalf("out of range: %d", v)
		}
	}
}

func TestInt64nRange(t *testing.T) {
	const n int64 = 1234567
	for i := 0; i < 200; i++ {
		v, err := Int64n(n)
		if err != nil {
			t.Fatalf("Int64n error: %v", err)
		}
		if v < 0 || v >= n {
			t.Fatalf("out of range: %d", v)
		}
	}
}

func TestFloat64Range(t *testing.T) {
	for i := 0; i < 1000; i++ {
		f, err := Float64()
		if err != nil {
			t.Fatalf("Float64 error: %v", err)
		}
		if !(f >= 0.0 && f < 1.0) {
			t.Fatalf("out of range: %f", f)
		}
	}
}

func TestUint64Deterministic(t *testing.T) {
	src := testutil.NewSeqReader(testutil.Uint64Bytes(0x0706050403020100))
	testutil.WithSource(t, src)
	v, err := Uint64()
	if err != nil {
		t.Fatalf("Uint64 returned error: %v", err)
	}
	const want = 0x0706050403020100
	if v != want {
		t.Fatalf("Uint64 = %x want %x", v, want)
	}
}

func TestUint64nRejectsHighValues(t *testing.T) {
	high := testutil.Uint64Bytes(^uint64(0))
	zero := testutil.Uint64Bytes(0)
	src := testutil.NewSeqReader(high, zero)
	testutil.WithSource(t, src)
	v, err := Uint64n(3)
	if err != nil {
		t.Fatalf("Uint64n returned error: %v", err)
	}
	if v != 0 {
		t.Fatalf("Uint64n result = %d want 0 after rejection", v)
	}
}

func TestUint64nZeroError(t *testing.T) {
	testutil.WithSource(t, testutil.NewSeqReader())
	if _, err := Uint64n(0); err != core.ErrInvalidRange {
		t.Fatalf("Uint64n error = %v want %v", err, core.ErrInvalidRange)
	}
}

func TestIntnInvalidParam(t *testing.T) {
	if _, err := Intn(0); err != core.ErrInvalidN {
		t.Fatalf("Intn error = %v want %v", err, core.ErrInvalidN)
	}
}

func TestFloat64Deterministic(t *testing.T) {
	src := testutil.NewSeqReader(testutil.Float64Bytes(0.25))
	testutil.WithSource(t, src)
	v, err := Float64()
	if err != nil {
		t.Fatalf("Float64 error: %v", err)
	}
	if math.Abs(v-0.25) > 1e-12 {
		t.Fatalf("Float64 = %f want 0.25", v)
	}
}

func TestMustWrappersPanicOnError(t *testing.T) {
	errSrc := testutil.ErrReader{Err: io.ErrUnexpectedEOF}
	testutil.WithSource(t, errSrc)
	expectsPanic := []struct {
		name string
		fn   func()
	}{
		{"MustUint64", func() { MustUint64() }},
		{"MustUint64n", func() { MustUint64n(10) }},
		{"MustIntn", func() { MustIntn(10) }},
		{"MustInt64n", func() { MustInt64n(10) }},
		{"MustFloat64", func() { MustFloat64() }},
	}
	for _, tc := range expectsPanic {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatalf("%s did not panic", tc.name)
				}
			}()
			tc.fn()
		})
	}
}
