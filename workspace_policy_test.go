//go:build randutil_policy
// +build randutil_policy

package randutil

import (
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/core"
)

func TestPolicyDisablesDeterministicRoot(t *testing.T) {
	ws := NewWorkspace(DeterministicRoot([]byte("seed")))
	if _, err := ws.Stream("alpha"); !errors.Is(err, core.ErrDeterministicDisabled) {
		t.Fatalf("Stream error = %v, want ErrDeterministicDisabled", err)
	}
	if err := ws.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
}
