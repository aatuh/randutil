//go:build randutil_must
// +build randutil_must

package uuid

import (
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestMustV4PanicsOnEntropyFailure(t *testing.T) {
	gen := New(core.New(testutil.ErrReader{Err: errors.New("entropy failure")}))
	defer func() {
		if recover() == nil {
			t.Fatalf("MustV4 did not panic on entropy failure")
		}
	}()
	gen.MustV4()
}

func TestMustParse(t *testing.T) {
	if MustParse("00000000-0000-0000-0000-000000000000").IsNil() != true {
		t.Fatalf("MustParse nil UUID not detected")
	}
	defer func() {
		if recover() == nil {
			t.Fatalf("MustParse did not panic on invalid input")
		}
	}()
	MustParse("not-a-uuid")
}
