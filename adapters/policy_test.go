//go:build randutil_policy
// +build randutil_policy

package adapters

import (
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/core"
)

func TestPolicyDisablesDeterministicSources(t *testing.T) {
	if _, err := DeterministicSource([]byte("seed")); !errors.Is(err, core.ErrDeterministicDisabled) {
		t.Fatalf("DeterministicSource error = %v, want ErrDeterministicDisabled", err)
	}
	if _, err := DeterministicSourceWithLabel([]byte("seed"), "label"); !errors.Is(err, core.ErrDeterministicDisabled) {
		t.Fatalf("DeterministicSourceWithLabel error = %v, want ErrDeterministicDisabled", err)
	}
}
