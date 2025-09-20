package uuid

import (
	"strings"
	"testing"
)

// Run with: go test -run=^$ -fuzz=FuzzParse -fuzztime=5s
func FuzzParse(f *testing.F) {
	f.Add("00000000-0000-0000-0000-000000000000")
	f.Add("12345678-1234-5678-9abc-def012345678")
	f.Fuzz(func(t *testing.T, input string) {
		u, err := Parse(input)
		if err != nil {
			return
		}
		if len(u) != 36 {
			t.Fatalf("parsed UUID length %d want 36", len(u))
		}
		s := string(u)
		if s != strings.ToLower(input) {
			t.Fatalf("Parse did not lowercase input: %s", s)
		}
		if _, err := u.Bytes(); err != nil {
			t.Fatalf("Bytes failed for parsed UUID: %v", err)
		}
	})
}
