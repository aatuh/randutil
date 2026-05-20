package ulid

import (
	"strings"
	"testing"
)

func FuzzParse(f *testing.F) {
	f.Add("01ARZ3NDEKTSV4RRFFQ69G5FAV")
	f.Add("00000000000000000000000000")
	f.Add("7ZZZZZZZZZZZZZZZZZZZZZZZZZ")
	f.Fuzz(func(t *testing.T, s string) {
		id, err := Parse(s)
		if err != nil {
			return
		}
		if len(id) != encodedLen {
			t.Fatalf("len=%d want %d", len(id), encodedLen)
		}
		if _, err := ulidEncoding.DecodeString(string(id)); err != nil {
			t.Fatalf("decode error: %v", err)
		}
		lower := strings.ToLower(string(id))
		parsed, err := Parse(lower)
		if err != nil {
			t.Fatalf("Parse error: %v", err)
		}
		if parsed != id {
			t.Fatalf("roundtrip mismatch: %s vs %s", parsed, id)
		}
	})
}
