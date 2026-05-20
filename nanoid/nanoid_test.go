package nanoid

import (
	"errors"
	"strings"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestIDLength(t *testing.T) {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	id, err := gen.ID(DefaultLength)
	if err != nil {
		t.Fatalf("ID error: %v", err)
	}
	if len(id) != DefaultLength {
		t.Fatalf("ID length=%d want %d", len(id), DefaultLength)
	}
}

func TestIDAlphabet(t *testing.T) {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	id, err := gen.ID(32)
	if err != nil {
		t.Fatalf("ID error: %v", err)
	}
	for _, r := range id {
		if !strings.ContainsRune(DefaultAlphabet, r) {
			t.Fatalf("unexpected rune %q in id", r)
		}
	}
}

func TestIDDeterministic(t *testing.T) {
	seed := []byte("seed")
	src1, err := adapters.DeterministicSource(seed)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	src2, err := adapters.DeterministicSource(seed)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	g1 := New(core.New(src1))
	g2 := New(core.New(src2))
	id1, err := g1.ID(16)
	if err != nil {
		t.Fatalf("ID error: %v", err)
	}
	id2, err := g2.ID(16)
	if err != nil {
		t.Fatalf("ID error: %v", err)
	}
	if id1 != id2 {
		t.Fatalf("deterministic mismatch: %s vs %s", id1, id2)
	}
}

func TestParse(t *testing.T) {
	id := "0123456789abcdef"
	parsed, err := Parse(id)
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if parsed != id {
		t.Fatalf("Parse mismatch: %s vs %s", parsed, id)
	}
	if _, err := Parse("invalid!"); err == nil {
		t.Fatalf("expected error for invalid id")
	}
}
