package adapters

import (
	"bytes"
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestFastSourceWithSourceDeterministic(t *testing.T) {
	seed := make([]byte, 32)
	for i := range seed {
		seed[i] = byte(i)
	}
	src1 := testutil.NewSeqReader(seed)
	src2 := testutil.NewSeqReader(seed)
	fast1, err := FastSourceWithSource(src1)
	if err != nil {
		t.Fatalf("FastSourceWithSource error: %v", err)
	}
	fast2, err := FastSourceWithSource(src2)
	if err != nil {
		t.Fatalf("FastSourceWithSource error: %v", err)
	}
	buf1 := make([]byte, 16)
	buf2 := make([]byte, 16)
	if _, err := io.ReadFull(fast1, buf1); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	if _, err := io.ReadFull(fast2, buf2); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	if !bytes.Equal(buf1, buf2) {
		t.Fatalf("fast sources mismatch")
	}
}
