package ulid

import (
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func BenchmarkULID(b *testing.B) {
	src, err := adapters.DeterministicSource([]byte("bench"))
	if err != nil {
		b.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.ULID()
	}
}
