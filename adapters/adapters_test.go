package adapters

import (
	"bytes"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestDeterministicSourceConsistency(t *testing.T) {
	seed := []byte("seed")
	s1 := DeterministicSource(seed)
	s2 := DeterministicSource(seed)
	buf1 := make([]byte, 32)
	buf2 := make([]byte, 32)
	if _, err := s1.Read(buf1); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if _, err := s2.Read(buf2); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if !bytes.Equal(buf1, buf2) {
		t.Fatalf("deterministic sources mismatch")
	}

	s3 := DeterministicSource([]byte("other"))
	buf3 := make([]byte, 32)
	if _, err := s3.Read(buf3); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if bytes.Equal(buf1, buf3) {
		t.Fatalf("different seeds produced identical output")
	}
}

func TestDeterministicSourceWithLabel(t *testing.T) {
	seed := []byte("seed")
	a1 := DeterministicSourceWithLabel(seed, "alpha")
	a2 := DeterministicSourceWithLabel(seed, "alpha")
	b1 := DeterministicSourceWithLabel(seed, "beta")

	bufA1 := make([]byte, 32)
	bufA2 := make([]byte, 32)
	bufB1 := make([]byte, 32)
	if _, err := a1.Read(bufA1); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if _, err := a2.Read(bufA2); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if _, err := b1.Read(bufB1); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	if !bytes.Equal(bufA1, bufA2) {
		t.Fatalf("same label produced different output")
	}
	if bytes.Equal(bufA1, bufB1) {
		t.Fatalf("different labels produced identical output")
	}
}

func TestDeterministicSourceConstructorDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("DeterministicSource panicked: %v", r)
		}
	}()
	src := DeterministicSource([]byte("seed"))
	buf := make([]byte, 8)
	if _, err := src.Read(buf); err != nil {
		t.Fatalf("Read error: %v", err)
	}
}

func TestLockedSourceWraps(t *testing.T) {
	src := LockedSource(testutil.NewSeqReader([]byte{9, 8}))
	buf := make([]byte, 3)
	if _, err := src.Read(buf); err != nil {
		t.Fatalf("Read error: %v", err)
	}
	want := []byte{9, 8, 8}
	if !bytes.Equal(buf, want) {
		t.Fatalf("buf = %v want %v", buf, want)
	}
}
