package dist

import (
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestZipfDeterministic(t *testing.T) {
	seed := []byte("zipf-seed")
	g1 := New(core.New(adapters.DeterministicSource(seed)))
	z1, err := g1.Zipf(1.2, 1.0, 10)
	if err != nil {
		t.Fatalf("Zipf error: %v", err)
	}
	g2 := New(core.New(adapters.DeterministicSource(seed)))
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
