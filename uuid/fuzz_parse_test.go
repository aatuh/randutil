package uuid

import (
	"strings"
	"testing"
)

// Run with: go test -run=^$ -fuzz=FuzzParse -fuzztime=5s
func FuzzParse(f *testing.F) {
	f.Add("00000000-0000-0000-0000-000000000000")
	f.Add("018cc820-d888-7a3b-ae5c-d5bc8d457263")
	f.Add("12345678-1234-5678-9abc-def012345678")
	f.Add("FFFFFFFF-FFFF-FFFF-FFFF-FFFFFFFFFFFF")
	f.Fuzz(func(t *testing.T, input string) {
		u, err := Parse(input)
		if err != nil {
			return
		}
		if len(u) != canonicalLen {
			t.Fatalf("parsed UUID length %d want %d", len(u), canonicalLen)
		}
		s := string(u)
		if s != strings.ToLower(input) {
			t.Fatalf("Parse did not lowercase input: %s", s)
		}
		if !isCanonicalUUID(u.String(), false) {
			t.Fatalf("non-canonical parse: %q", u)
		}
		b, err := u.Bytes()
		if err != nil {
			t.Fatalf("Bytes failed for parsed UUID: %v", err)
		}
		if roundtrip := fromBytes(b); roundtrip != u {
			t.Fatalf("roundtrip mismatch: %s vs %s", roundtrip, u)
		}
	})
}
