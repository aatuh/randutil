package numeric

import "testing"

func TestBoolReturnsValue(t *testing.T) {
	b, err := Bool()
	if err != nil {
		t.Fatalf("Bool error: %v", err)
	}
	if b != true && b != false {
		t.Fatalf("Bool produced non-boolean value: %v", b)
	}
}
