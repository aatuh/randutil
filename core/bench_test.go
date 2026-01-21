package core

import "testing"

type seqSource struct {
	v byte
}

func (s *seqSource) Read(p []byte) (int, error) {
	for i := range p {
		p[i] = s.v
		s.v++
	}
	return len(p), nil
}

func BenchmarkUint64n(b *testing.B) {
	gen := New(&seqSource{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Uint64n(1 << 32)
	}
}

func BenchmarkFloat64(b *testing.B) {
	gen := New(&seqSource{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Float64()
	}
}

func BenchmarkBytes16(b *testing.B) {
	gen := New(&seqSource{})
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = gen.Bytes(16)
	}
}
