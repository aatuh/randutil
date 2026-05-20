package adapters

import (
	"bytes"
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestRecorderReplay(t *testing.T) {
	src := testutil.NewSeqReader([]byte{1, 2, 3})
	rec := NewRecorder(src)
	buf := make([]byte, 4)
	if _, err := io.ReadFull(rec, buf); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	recorded := rec.Bytes()
	if !bytes.Equal(buf, recorded) {
		t.Fatalf("recorded=%v want %v", recorded, buf)
	}
	replay := ReplaySource(recorded)
	got := make([]byte, len(recorded))
	if _, err := io.ReadFull(replay, got); err != nil {
		t.Fatalf("Replay ReadFull error: %v", err)
	}
	if !bytes.Equal(got, recorded) {
		t.Fatalf("replay=%v want %v", got, recorded)
	}
	rec.Reset()
	if len(rec.Bytes()) != 0 {
		t.Fatalf("expected Reset to clear recording")
	}
}

func TestReplaySourceShortRead(t *testing.T) {
	replay := ReplaySource([]byte{9, 8, 7})
	buf := make([]byte, 4)
	n, err := replay.Read(buf)
	if err == nil {
		t.Fatalf("expected error on short read")
	}
	if n != 3 {
		t.Fatalf("short read n=%d want 3", n)
	}
}
