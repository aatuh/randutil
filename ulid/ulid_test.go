package ulid

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestULIDLengthAndAlphabet(t *testing.T) {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	id, err := gen.ULID()
	if err != nil {
		t.Fatalf("ULID error: %v", err)
	}
	if len(id) != encodedLen {
		t.Fatalf("ULID length=%d want %d", len(id), encodedLen)
	}
	for _, r := range id {
		if !strings.ContainsRune(crockfordAlphabet, r) {
			t.Fatalf("unexpected rune %q in ULID", r)
		}
	}
}

func TestULIDTimestamp(t *testing.T) {
	fixed := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := NewWithClock(
		core.New(src),
		func() time.Time { return fixed },
	)
	id, err := gen.ULID()
	if err != nil {
		t.Fatalf("ULID error: %v", err)
	}
	raw, err := ulidEncoding.DecodeString(string(id))
	if err != nil {
		t.Fatalf("decode error: %v", err)
	}
	if len(raw) != 16 {
		t.Fatalf("decoded length=%d want 16", len(raw))
	}
	ms := int64(raw[0])<<40 | int64(raw[1])<<32 | int64(raw[2])<<24 |
		int64(raw[3])<<16 | int64(raw[4])<<8 | int64(raw[5])
	if ms != fixed.UnixMilli() {
		t.Fatalf("timestamp=%d want %d", ms, fixed.UnixMilli())
	}
}

func TestParse(t *testing.T) {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	id, err := gen.ULID()
	if err != nil {
		t.Fatalf("ULID error: %v", err)
	}
	lower := strings.ToLower(string(id))
	parsed, err := Parse(lower)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if parsed != id {
		t.Fatalf("Parse mismatch: %s vs %s", parsed, id)
	}
	if _, err := Parse("invalid"); err == nil {
		t.Fatalf("expected error for invalid ULID")
	}
}
