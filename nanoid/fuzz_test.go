package nanoid

import "testing"

func FuzzParse(f *testing.F) {
	f.Add("0123456789abcdefghij")
	f.Add("_-0123456789abcdefghijklmnopqrstuvwxyz")
	f.Fuzz(func(t *testing.T, s string) {
		id, err := Parse(s)
		if err != nil {
			return
		}
		if _, err := ParseWithAlphabet(id, DefaultAlphabet); err != nil {
			t.Fatalf("ParseWithAlphabet error: %v", err)
		}
	})
}
