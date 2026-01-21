package randstring

import (
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestStringWithCharsetRejectsOutOfRangeBytes(t *testing.T) {
	gen := New(core.New(testutil.NewSeqReader([]byte{255, 1, 2, 3})))
	s, err := gen.StringWithCharset(3, "abcde")
	if err != nil {
		t.Fatalf("StringWithCharset error: %v", err)
	}
	if s != "bcd" {
		t.Fatalf("StringWithCharset = %q want %q", s, "bcd")
	}
}
