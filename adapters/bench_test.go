package adapters

import "testing"

func BenchmarkCryptoSourceRead(b *testing.B) {
	src := CryptoSource()
	buf := make([]byte, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = src.Read(buf)
	}
}

func BenchmarkDerivedSourceRead(b *testing.B) {
	src, err := DeriveSource([]byte("bench"), "stream")
	if err != nil {
		b.Fatalf("DeriveSource error: %v", err)
	}
	buf := make([]byte, 1024)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = src.Read(buf)
	}
}
