//go:build randutil_must
// +build randutil_must

package nanoid

import (
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestMustIDPanics(t *testing.T) {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		if errors.Is(err, core.ErrDeterministicDisabled) {
			t.Skip("deterministic sources disabled")
		}
		t.Fatalf("DeterministicSource error: %v", err)
	}
	gen := New(core.New(src))
	defer func() {
		if recover() == nil {
			t.Fatalf("MustID did not panic on invalid length")
		}
	}()
	gen.MustID(-1)
}
