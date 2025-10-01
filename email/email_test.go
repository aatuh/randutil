package email

import (
	"strings"
	"testing"

	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestEmailIgnoresTotalLengthWhenOptionsProvided(t *testing.T) {
	// TLD override present -> ignore TotalLength.
	e, err := Email(EmailOptions{TLD: "org", TotalLength: 9})
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	if !strings.HasSuffix(e, ".org") {
		t.Fatalf("expected .org: %s", e)
	}
	if len(e) == 9 {
		t.Fatalf("unexpected exact TotalLength honored with TLD override")
	}

	// LocalPart override present -> ignore TotalLength.
	e, err = Email(EmailOptions{LocalPart: "aaa", TotalLength: 50})
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	if !strings.HasPrefix(e, "aaa@") {
		t.Fatalf("expected local override: %s", e)
	}
	if len(e) == 50 {
		t.Fatalf("unexpected TotalLength honored with local override")
	}

	// DomainPart override present -> ignore TotalLength.
	e, err = Email(EmailOptions{DomainPart: "example", TotalLength: 8})
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	if !strings.Contains(e, "@example.") {
		t.Fatalf("expected domain override: %s", e)
	}
	if len(e) == 8 {
		t.Fatalf("unexpected TotalLength honored with domain override")
	}
}

func TestEmailSimple(t *testing.T) {
	// Test with valid length
	email, err := EmailSimple(10)
	if err != nil {
		t.Fatalf("EmailSimple error: %v", err)
	}
	if len(email) != 10 {
		t.Fatalf("Expected length 10, got %d", len(email))
	}
	if !strings.Contains(email, "@") {
		t.Fatalf("Email should contain @: %s", email)
	}
	if !strings.HasSuffix(email, ".com") {
		t.Fatalf("Email should end with .com: %s", email)
	}

	// Test with too small length
	_, err = EmailSimple(5)
	if err == nil {
		t.Fatal("Expected error for too small length")
	}
}

func TestEmailWithCustomLocal(t *testing.T) {
	email, err := EmailWithCustomLocal("testuser")
	if err != nil {
		t.Fatalf("EmailWithCustomLocal error: %v", err)
	}
	if !strings.HasPrefix(email, "testuser@") {
		t.Fatalf("Expected email to start with testuser@: %s", email)
	}
}

func TestEmailWithCustomDomain(t *testing.T) {
	email, err := EmailWithCustomDomain("example")
	if err != nil {
		t.Fatalf("EmailWithCustomDomain error: %v", err)
	}
	if !strings.Contains(email, "@example.") {
		t.Fatalf("Expected email to contain @example.: %s", email)
	}
}

func TestEmailWithCustomTLD(t *testing.T) {
	email, err := EmailWithCustomTLD("org")
	if err != nil {
		t.Fatalf("EmailWithCustomTLD error: %v", err)
	}
	if !strings.HasSuffix(email, ".org") {
		t.Fatalf("Expected email to end with .org: %s", email)
	}
}

func TestEmailWithRandomTLD(t *testing.T) {
	email, err := EmailWithRandomTLD()
	if err != nil {
		t.Fatalf("EmailWithRandomTLD error: %v", err)
	}
	// Should contain @ and have a TLD
	if !strings.Contains(email, "@") {
		t.Fatalf("Email should contain @: %s", email)
	}
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		t.Fatalf("Email should have exactly one @: %s", email)
	}
	domain := parts[1]
	if !strings.Contains(domain, ".") {
		t.Fatalf("Domain should contain a dot: %s", domain)
	}
}

func TestEmailWithoutTLD(t *testing.T) {
	email, err := EmailWithoutTLD()
	if err != nil {
		t.Fatalf("EmailWithoutTLD error: %v", err)
	}
	if !strings.Contains(email, "@") {
		t.Fatalf("Email should contain @: %s", email)
	}
	if strings.Contains(email, ".") {
		t.Fatalf("Email should not contain dots (no TLD): %s", email)
	}
}

func TestEmailGeneratorUsesOwnSource(t *testing.T) {
	// Fixed stream will produce deterministic output.
	src := testutil.NewSeqReader([]byte("abcdef"))
	g := New(src)

	e1, err := g.Email(EmailOptions{})
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	e2, err := g.Email(EmailOptions{})
	if err != nil {
		t.Fatalf("Email error: %v", err)
	}
	if e1 == "" || e2 == "" || e1 == e2 {
		t.Fatalf("expected two different, deterministic emails from same source")
	}
}
