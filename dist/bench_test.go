package dist

import (
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func BenchmarkNormal(b *testing.B) {
	src, err := adapters.DeterministicSource([]byte("bench"))
	if err != nil {
		b.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Normal(0, 1)
	}
}

func BenchmarkPoisson(b *testing.B) {
	src, err := adapters.DeterministicSource([]byte("bench"))
	if err != nil {
		b.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Poisson(12)
	}
}

func BenchmarkGamma(b *testing.B) {
	src, err := adapters.DeterministicSource([]byte("bench"))
	if err != nil {
		b.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Gamma(2.5, 1.3)
	}
}
