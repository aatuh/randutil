package dist

import (
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestZipfDeterministic(t *testing.T) {
	seed := []byte("zipf-seed")
	src1, err := adapters.DeterministicSource(seed)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	g1 := New(core.New(src1))
	z1, err := g1.Zipf(1.2, 1.0, 10)
	if err != nil {
		t.Fatalf("Zipf error: %v", err)
	}
	src2, err := adapters.DeterministicSource(seed)
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	g2 := New(core.New(src2))
	z2, err := g2.Zipf(1.2, 1.0, 10)
	if err != nil {
		t.Fatalf("Zipf error: %v", err)
	}
	for i := 0; i < 5; i++ {
		v1, err := z1.Next()
		if err != nil {
			t.Fatalf("Zipf Next error: %v", err)
		}
		v2, err := z2.Next()
		if err != nil {
			t.Fatalf("Zipf Next error: %v", err)
		}
		if v1 != v2 {
			t.Fatalf("Zipf mismatch at %d: %d vs %d", i, v1, v2)
		}
	}
}
