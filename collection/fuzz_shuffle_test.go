package collection

import "testing"

func FuzzShuffle(f *testing.F) {
	f.Add(0)
	f.Add(1)
	f.Add(10)
	f.Add(100)
	f.Fuzz(func(t *testing.T, n int) {
		if n < 0 {
			return
		}
		s := make([]int, n)
		for i := range s {
			s[i] = i
		}
		if err := Shuffle(s); err != nil {
			t.Fatalf("Shuffle error: %v", err)
		}
		// Permutation check.
		seen := make([]bool, n)
		for _, v := range s {
			if v < 0 || v >= n {
				t.Fatalf("value out of range: %d", v)
			}
			if seen[v] {
				t.Fatalf("duplicate value: %d", v)
			}
			seen[v] = true
		}
	})
}
