//go:build randutil_must
// +build randutil_must

package collection

import "testing"

func TestMustChoice(t *testing.T) {
	if MustChoice("a", "b") == "" {
		t.Fatal("MustChoice returned empty result")
	}
}

func TestMustPickPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustPickOne did not panic for empty slice")
		}
	}()
	MustPickOne([]int{})
}
