package numeric

import (
	"errors"
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestBytesAndFill(t *testing.T) {
	src := testutil.NewSeqReader([]byte{9, 8, 7, 6, 5})
	gen := New(core.New(src))
	data, err := gen.Bytes(4)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	want := []byte{9, 8, 7, 6}
	if !equalBytes(data, want) {
		t.Fatalf("Bytes = %v want %v", data, want)
	}
	buf := make([]byte, 3)
	if err := gen.Fill(buf); err != nil {
		t.Fatalf("Fill error: %v", err)
	}
	expected := []byte{5, 5, 5}
	if !equalBytes(buf, expected) {
		t.Fatalf("Fill buf=%v want %v", buf, expected)
	}
}

func TestBytesInvalidLength(t *testing.T) {
	gen := New(core.New(testutil.NewSeqReader()))
	if _, err := gen.Bytes(-1); !errors.Is(err, core.ErrNegativeLength) {
		t.Fatalf("Bytes(-1) error = %v want %v", err, core.ErrNegativeLength)
	}
}

func TestMustBytesPanicOnError(t *testing.T) {
	errSrc := testutil.ErrReader{Err: io.ErrUnexpectedEOF}
	gen := New(core.New(errSrc))
	defer func() {
		if recover() == nil {
			t.Fatalf("MustBytes did not panic")
		}
	}()
	gen.MustBytes(1)
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
