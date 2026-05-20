package adapters

import (
	"io"
	"sync/atomic"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestCountingSourceCounts(t *testing.T) {
	src := testutil.NewSeqReader([]byte{1, 2, 3})
	counter := NewCountingSource(src, nil)
	buf := make([]byte, 5)
	if _, err := io.ReadFull(counter, buf); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	if got := counter.Count(); got != 5 {
		t.Fatalf("Count=%d want 5", got)
	}
}

func TestCountingSourceHook(t *testing.T) {
	src := testutil.NewSeqReader([]byte{1, 2, 3})
	var total atomic.Uint64
	counter := NewCountingSource(src, func(delta uint64) {
		total.Add(delta)
	})
	buf := make([]byte, 4)
	if _, err := io.ReadFull(counter, buf); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	if got := total.Load(); got != 4 {
		t.Fatalf("Hook total=%d want 4", got)
	}
}
