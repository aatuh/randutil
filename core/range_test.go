package core

import (
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestInt64RangeFullRangeCanReturnMinInt64(t *testing.T) {
	gen := New(testutil.NewSeqReader(testutil.Uint64Bytes(0)))
	got, err := gen.Int64Range(minInt64, maxInt64)
	if err != nil {
		t.Fatalf("Int64Range full range returned error: %v", err)
	}
	if got != minInt64 {
		t.Fatalf("Int64Range full range = %d want %d", got, minInt64)
	}
}

func TestIntRangeFullRangeCanReturnMinInt(t *testing.T) {
	gen := New(testutil.NewSeqReader(testutil.Uint64Bytes(0)))
	got, err := gen.IntRange(minInt, maxInt)
	if err != nil {
		t.Fatalf("IntRange full range returned error: %v", err)
	}
	if got != minInt {
		t.Fatalf("IntRange full range = %d want %d", got, minInt)
	}
}
