//go:build randutil_must
// +build randutil_must

package numeric

import (
	"io"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestMustBytesPanicOnError(t *testing.T) {
	errSrc := testutil.ErrReader{Err: io.ErrUnexpectedEOF}
	gen := New(core.New(errSrc))
	defer func() {
		if recover() == nil {
			t.Fatalf("MustBytes did not panic")
		}
	}()
	gen.MustBytes(1)
}
