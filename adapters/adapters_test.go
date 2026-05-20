package adapters

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestDeterministicSourceConsistency(t *testing.T) {
	seed := []byte("seed")
	s1 := mustDeterministicSource(t, seed)
	s2 := mustDeterministicSource(t, seed)
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

	s3 := mustDeterministicSource(t, []byte("other"))
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
	a1 := mustDeterministicSourceWithLabel(t, seed, "alpha")
	a2 := mustDeterministicSourceWithLabel(t, seed, "alpha")
	b1 := mustDeterministicSourceWithLabel(t, seed, "beta")

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

func TestDeriveSourceWithLabel(t *testing.T) {
	seed := []byte("seed")
	a1 := mustDeriveSource(t, seed, "alpha")
	a2 := mustDeriveSource(t, seed, "alpha")
	b1 := mustDeriveSource(t, seed, "beta")

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

func TestChaChaSourceReturnsErrorWhenExhausted(t *testing.T) {
	var key [32]byte
	var nonce [12]byte
	src, err := newChaChaSourceWithLimit(key, nonce, 4)
	if err != nil {
		t.Fatalf("newChaChaSourceWithLimit error: %v", err)
	}

	buf := make([]byte, 4)
	if n, err := src.Read(buf); n != 4 || err != nil {
		t.Fatalf("first Read = (%d, %v), want (4, nil)", n, err)
	}

	buf = []byte{1, 2, 3}
	n, err := src.Read(buf)
	if n != 0 || !errors.Is(err, core.ErrSourceExhausted) {
		t.Fatalf("exhausted Read = (%d, %v), want (0, ErrSourceExhausted)", n, err)
	}
	if !bytes.Equal(buf, []byte{0, 0, 0}) {
		t.Fatalf("exhausted Read did not zero buffer: %v", buf)
	}
}

func TestChaChaSourceRejectsOversizedRead(t *testing.T) {
	var key [32]byte
	var nonce [12]byte
	src, err := newChaChaSourceWithLimit(key, nonce, 4)
	if err != nil {
		t.Fatalf("newChaChaSourceWithLimit error: %v", err)
	}

	buf := []byte{1, 2, 3, 4, 5}
	n, err := src.Read(buf)
	if n != 0 || !errors.Is(err, core.ErrSourceExhausted) {
		t.Fatalf("oversized Read = (%d, %v), want (0, ErrSourceExhausted)", n, err)
	}
	if !bytes.Equal(buf, []byte{0, 0, 0, 0, 0}) {
		t.Fatalf("oversized Read did not zero buffer: %v", buf)
	}
}

func TestDeterministicSourceConstructorDoesNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Fatalf("DeterministicSource panicked: %v", r)
		}
	}()
	src := mustDeterministicSource(t, []byte("seed"))
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

func TestBufferedSourceMatchesUnderlying(t *testing.T) {
	data := []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	src := testutil.NewSeqReader(data)
	wantSrc := testutil.NewSeqReader(data)
	bufSrc := BufferedSourceWithSize(src, 3)

	sizes := []int{1, 2, 5, 9, 4, 7}
	for _, size := range sizes {
		got := make([]byte, size)
		want := make([]byte, size)
		if _, err := io.ReadFull(bufSrc, got); err != nil {
			t.Fatalf("buffered Read error: %v", err)
		}
		if _, err := io.ReadFull(wantSrc, want); err != nil {
			t.Fatalf("direct Read error: %v", err)
		}
		if !bytes.Equal(got, want) {
			t.Fatalf("buffered output mismatch size %d: %v vs %v", size, got, want)
		}
	}
}

func mustDeterministicSource(t testing.TB, seed []byte) core.Source {
	t.Helper()
	src, err := DeterministicSource(seed)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	return src
}

func mustDeterministicSourceWithLabel(t testing.TB, seed []byte, label string) core.Source {
	t.Helper()
	src, err := DeterministicSourceWithLabel(seed, label)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSourceWithLabel error: %v", err)
	}
	return src
}

func mustDeriveSource(t testing.TB, seed []byte, label string) core.Source {
	t.Helper()
	src, err := DeriveSource(seed, label)
	if err != nil {
		t.Fatalf("DeriveSource error: %v", err)
	}
	return src
}
