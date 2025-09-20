package string

import (
	"strings"
	"testing"
)

func TestEmailStructure(t *testing.T) {
	email, err := Email(12)
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	if len(email) != 12 {
		t.Fatalf("Email length = %d want 12", len(email))
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		t.Fatalf("Email missing @: %s", email)
	}
	domain := parts[1]
	if !strings.HasSuffix(domain, ".com") {
		t.Fatalf("Email domain missing .com: %s", email)
	}
	if len(parts[0]) == 0 || len(domain) <= 4 {
		t.Fatalf("Email sides too short: %s", email)
	}
}

func TestEmailErrors(t *testing.T) {
	if _, err := Email(6); err == nil {
		t.Fatalf("expected error for too-short email")
	}
}
