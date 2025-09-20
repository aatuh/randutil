package collection

import "testing"

// Run with: go test -run=^$ -fuzz=FuzzWeightedChoice -fuzztime=5s
func FuzzWeightedChoice(f *testing.F) {
	f.Add([]byte{1, 2, 3}, []byte{1, 1, 1})
	f.Add([]byte{9, 8}, []byte{0, 5})
	f.Fuzz(func(t *testing.T, itemsSeed []byte, weightsSeed []byte) {
		if len(itemsSeed) == 0 {
			if _, err := WeightedChoice([]byte{}, []float64{}); err == nil {
				t.Fatalf("expected error on empty input")
			}
			return
		}
		items := make([]byte, len(itemsSeed))
		copy(items, itemsSeed)

		weights := make([]float64, len(itemsSeed))
		if len(weightsSeed) == 0 {
			if _, err := WeightedChoice(items, weights); err == nil {
				t.Fatalf("expected error when weights missing")
			}
			return
		}
		sum := 0.0
		for i := range weights {
			w := float64(int(weightsSeed[i%len(weightsSeed)]))
			weights[i] = w
			sum += w
		}
		if sum <= 0 {
			if _, err := WeightedChoice(items, weights); err == nil {
				t.Fatalf("expected error with non-positive sum")
			}
			return
		}
		v, err := WeightedChoice(items, weights)
		if err != nil {
			t.Fatalf("WeightedChoice error: %v", err)
		}
		found := false
		for i, item := range items {
			if weights[i] > 0 && item == v {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("WeightedChoice returned %v not tied to positive weight", v)
		}
	})
}
