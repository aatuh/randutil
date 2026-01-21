package adapters

import (
	"bytes"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestLockedRNGDelegates(t *testing.T) {
	rng := core.New(testutil.NewSeqReader([]byte{7, 8}))
	locked := LockedRNG(rng)
	b, err := locked.Bytes(3)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	want := []byte{7, 8, 8}
	if !bytes.Equal(b, want) {
		t.Fatalf("bytes = %v want %v", b, want)
	}
}
