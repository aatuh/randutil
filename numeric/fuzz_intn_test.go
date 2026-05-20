package numeric

import (
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

// Run with: go test -run=^$ -fuzz=Fuzz -fuzztime=5s
func FuzzUint64n(f *testing.F) {
	f.Add(uint64(1))
	f.Add(uint64(2))
	f.Add(uint64(10))
	f.Add(^uint64(0))
	f.Fuzz(func(t *testing.T, n uint64) {
		if n == 0 {
			_, err := Uint64n(n)
			if err == nil {
				t.Fatalf("expected error for n=0")
			}
			return
		}
		gen := New(core.New(testutil.NewSeqReader(testutil.Uint64Bytes(0))))
		v, err := gen.Uint64n(n)
		if err != nil {
			t.Fatalf("Uint64n error: %v", err)
		}
		if v >= n {
			t.Fatalf("out of range: %d >= %d", v, n)
		}
	})
}
