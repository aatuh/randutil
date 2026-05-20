//go:build randutil_must
// +build randutil_must

package numeric

import (
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestMustWrappersPanicOnError(t *testing.T) {
	errSrc := testutil.ErrReader{Err: io.ErrUnexpectedEOF}
	gen := New(core.New(errSrc))
	expectsPanic := []struct {
		name string
		fn   func()
	}{
		{"MustUint64", func() { gen.MustUint64() }},
		{"MustUint64n", func() { gen.MustUint64n(10) }},
		{"MustIntn", func() { gen.MustIntn(10) }},
		{"MustInt64n", func() { gen.MustInt64n(10) }},
		{"MustFloat64", func() { gen.MustFloat64() }},
	}
	for _, tc := range expectsPanic {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				if recover() == nil {
					t.Fatalf("%s did not panic", tc.name)
				}
			}()
			tc.fn()
		})
	}
}
