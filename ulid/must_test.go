//go:build randutil_must
// +build randutil_must

package ulid

import "testing"

func TestMustParsePanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatalf("MustParse did not panic on invalid input")
		}
	}()
	MustParse("invalid")
}
