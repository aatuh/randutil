package randtime

import (
	"errors"
	"testing"
	"time"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestJitterDeterministic(t *testing.T) {
	src := testutil.NewSeqReader(testutil.Float64Bytes(0.75))
	gen := New(core.New(src))
	got, err := gen.Jitter(10*time.Second, 0.2)
	if err != nil {
		t.Fatalf("Jitter error: %v", err)
	}
	if got != 11*time.Second {
		t.Fatalf("Jitter=%s want %s", got, 11*time.Second)
	}
}

func TestJitterErrors(t *testing.T) {
	if _, err := Jitter(-time.Second, 0.1); !errors.Is(err, core.ErrNegativeDuration) {
		t.Fatalf("expected ErrNegativeDuration, got %v", err)
	}
	if _, err := Jitter(time.Second, -0.1); !errors.Is(err, core.ErrInvalidJitter) {
		t.Fatalf("expected ErrInvalidJitter, got %v", err)
	}
	if _, err := Jitter(time.Second, 1.5); !errors.Is(err, core.ErrInvalidJitter) {
		t.Fatalf("expected ErrInvalidJitter, got %v", err)
	}
}
