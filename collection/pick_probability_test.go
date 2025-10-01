package collection

import (
	"testing"
)

func TestPickByProbability(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}

	// Test with probability 0 (should return empty)
	result, err := PickByProbability(items, 0.0)
	if err != nil {
		t.Fatalf("PickByProbability error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result with probability 0, got %v", result)
	}

	// Test with probability 1 (should return all items)
	result, err = PickByProbability(items, 1.0)
	if err != nil {
		t.Fatalf("PickByProbability error: %v", err)
	}
	if len(result) != len(items) {
		t.Errorf("Expected all items with probability 1, got %d items", len(result))
	}

	// Test with probability 0.5 (should return some items)
	result, err = PickByProbability(items, 0.5)
	if err != nil {
		t.Fatalf("PickByProbability error: %v", err)
	}
	// Result should be between 0 and len(items)
	if len(result) > len(items) {
		t.Errorf("Invalid result length: %d", len(result))
	}
}

func TestPickByProbabilityEdgeCases(t *testing.T) {
	// Test empty slice
	result, err := PickByProbability([]string{}, 0.5)
	if err != nil {
		t.Fatalf("PickByProbability error: %v", err)
	}
	if len(result) != 0 {
		t.Errorf("Expected empty result for empty input, got %v", result)
	}

	// Test invalid probability (negative)
	_, err = PickByProbability([]string{"a"}, -0.1)
	if err == nil {
		t.Error("Expected error for negative probability")
	}

	// Test invalid probability (> 1)
	_, err = PickByProbability([]string{"a"}, 1.1)
	if err == nil {
		t.Error("Expected error for probability > 1")
	}
}
