package adapters

import (
	"encoding/hex"
	"errors"
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/core"
)

func TestDeterministicSourceGolden(t *testing.T) {
	src, err := DeterministicSource([]byte("seed"))
	if errors.Is(err, core.ErrDeterministicDisabled) {
		t.Skip("deterministic sources disabled")
	}
	if err != nil {
		t.Fatalf("DeterministicSource error: %v", err)
	}
	buf := make([]byte, 32)
	if _, err := io.ReadFull(src, buf); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	const wantHex = "cb262e257620ac804edbe10735c8417ab4ce48cc0ddf4bc08519f8212198fc8d"
	if hex.EncodeToString(buf) != wantHex {
		t.Fatalf("golden mismatch: %x", buf)
	}
}

func TestDeriveSourceGolden(t *testing.T) {
	src, err := DeriveSource([]byte("seed"), "alpha")
	if err != nil {
		t.Fatalf("DeriveSource error: %v", err)
	}
	buf := make([]byte, 32)
	if _, err := io.ReadFull(src, buf); err != nil {
		t.Fatalf("ReadFull error: %v", err)
	}
	const wantHex = "461b3c8d31c3ac27ac70f8271059f925ce151ab71e035e83cbe873ff9a60f6a5"
	if hex.EncodeToString(buf) != wantHex {
		t.Fatalf("golden mismatch: %x", buf)
	}
}
