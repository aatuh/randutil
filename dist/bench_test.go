package dist

import (
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func BenchmarkNormal(b *testing.B) {
	gen := New(core.New(adapters.DeterministicSource([]byte("bench"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Normal(0, 1)
	}
}

func BenchmarkPoisson(b *testing.B) {
	gen := New(core.New(adapters.DeterministicSource([]byte("bench"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Poisson(12)
	}
}

func BenchmarkGamma(b *testing.B) {
	gen := New(core.New(adapters.DeterministicSource([]byte("bench"))))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Gamma(2.5, 1.3)
	}
}
