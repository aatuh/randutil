package numeric

import (
	"errors"
	"io"
	"testing"

	"github.com/aatuh/randutil/core"
	"github.com/aatuh/randutil/internal/testutil"
)

func TestBytesAndFill(t *testing.T) {
	src := testutil.NewSeqReader([]byte{9, 8, 7, 6, 5})
	testutil.WithSource(t, src)
	data, err := Bytes(4)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	want := []byte{9, 8, 7, 6}
	if !equalBytes(data, want) {
		t.Fatalf("Bytes = %v want %v", data, want)
	}
	buf := make([]byte, 3)
	if err := Fill(buf); err != nil {
		t.Fatalf("Fill error: %v", err)
	}
	expected := []byte{5, 5, 5}
	if !equalBytes(buf, expected) {
		t.Fatalf("Fill buf=%v want %v", buf, expected)
	}
}

func TestBytesInvalidLength(t *testing.T) {
	if _, err := Bytes(-1); !errors.Is(err, core.ErrInvalidN) {
		t.Fatalf("Bytes(-1) error = %v want %v", err, core.ErrInvalidN)
	}
}

func TestMustBytesPanicOnError(t *testing.T) {
	errSrc := testutil.ErrReader{Err: io.ErrUnexpectedEOF}
	testutil.WithSource(t, errSrc)
	defer func() {
		if recover() == nil {
			t.Fatalf("MustBytes did not panic")
		}
	}()
	MustBytes(1)
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
