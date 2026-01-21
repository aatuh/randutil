package core

import (
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestGeneratorAsReader(t *testing.T) {
	// Test that core.Generator implements io.Reader
	var _ io.Reader = &Generator{}

	// Test that we can read from a generator
	gen := New(nil)
	buf := make([]byte, 10)

	n, err := gen.Read(buf)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	if n != 10 {
		t.Errorf("Expected to read 10 bytes, got %d", n)
	}

	// Verify we got some non-zero bytes (very unlikely to get all zeros)
	allZero := true
	for _, b := range buf {
		if b != 0 {
			allZero = false
			break
		}
	}
	if allZero {
		t.Error("Expected some non-zero bytes from random source")
	}
}

func TestFillZeroesOnError(t *testing.T) {
	gen := New(testutil.ErrReader{Err: io.ErrUnexpectedEOF})
	buf := []byte{1, 2, 3, 4}
	if err := gen.Fill(buf); err == nil {
		t.Fatalf("Fill unexpectedly succeeded")
	}
	for i, b := range buf {
		if b != 0 {
			t.Fatalf("buf[%d]=%d want 0", i, b)
		}
	}
}
