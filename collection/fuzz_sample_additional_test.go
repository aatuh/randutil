package collection

import "testing"

// Run with: go test -run=^$ -fuzz=FuzzSample -fuzztime=5s
func FuzzSample(f *testing.F) {
	f.Add(uint8(5), uint8(3))
	f.Add(uint8(10), uint8(0))
	f.Fuzz(func(t *testing.T, sizeSeed uint8, kSeed uint8) {
		size := int(sizeSeed % 64)
		s := make([]int, size)
		for i := range s {
			s[i] = i
		}
		k := int(kSeed % 8)
		if k > size {
			if _, err := Sample(s, k); err == nil {
				t.Fatalf("expected error when k>%d", size)
			}
			return
		}
		out, err := Sample(s, k)
		if err != nil {
			t.Fatalf("Sample error: %v", err)
		}
		if len(out) != k {
			t.Fatalf("Sample length %d want %d", len(out), k)
		}
		seen := make(map[int]struct{}, len(out))
		for _, v := range out {
			if v < 0 || v >= size {
				t.Fatalf("Sample produced %d outside range [0,%d)", v, size)
			}
			if _, dup := seen[v]; dup {
				t.Fatalf("Sample produced duplicate element %d", v)
			}
			seen[v] = struct{}{}
		}
	})
}
