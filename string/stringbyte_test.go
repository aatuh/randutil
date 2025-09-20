package string

import (
	"encoding/base64"
	"encoding/hex"
	"strings"
	"testing"
)

func TestGetWithCustomCharsetErrors(t *testing.T) {
	if _, err := GetWithCustomCharset(-1, "abc"); err == nil {
		t.Fatalf("expected error for negative length")
	}
	if _, err := GetWithCustomCharset(1, ""); err == nil {
		t.Fatalf("expected error for empty charset")
	}
}

func TestGetWithCustomCharsetUsesCharset(t *testing.T) {
	charset := "abc"
	s, err := GetWithCustomCharset(32, charset)
	if err != nil {
		t.Fatalf("GetWithCustomCharset error: %v", err)
	}
	if len(s) != 32 {
		t.Fatalf("length = %d want 32", len(s))
	}
	for _, ch := range s {
		if !strings.ContainsRune(charset, ch) {
			t.Fatalf("character %q not in charset", ch)
		}
	}
}

func TestStringUsesDefaultCharset(t *testing.T) {
	s, err := String(16)
	if err != nil {
		t.Fatalf("String error: %v", err)
	}
	if len(s) != 16 {
		t.Fatalf("String length = %d want 16", len(s))
	}
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	for _, ch := range s {
		if !strings.ContainsRune(charset, ch) {
			t.Fatalf("String produced invalid rune %q", ch)
		}
	}
}

func TestHexAndBase64(t *testing.T) {
	s, err := Hex(16)
	if err != nil {
		t.Fatalf("Hex error: %v", err)
	}
	if len(s) != 16 {
		t.Fatalf("Hex length = %d want 16", len(s))
	}
	if _, err := hex.DecodeString(s); err != nil {
		t.Fatalf("Hex produced invalid string: %v", err)
	}
	if _, err := Hex(5); err == nil {
		t.Fatalf("expected error for odd length Hex")
	}

	b64, err := Base64(12)
	if err != nil {
		t.Fatalf("Base64 error: %v", err)
	}
	if _, err := base64.StdEncoding.DecodeString(b64); err != nil {
		t.Fatalf("Base64 output invalid: %v", err)
	}
}

func TestStringSliceLenRange(t *testing.T) {
	slice, err := StringSlice(5, 3, 6)
	if err != nil {
		t.Fatalf("StringSlice error: %v", err)
	}
	if len(slice) != 5 {
		t.Fatalf("StringSlice len = %d want 5", len(slice))
	}
	for _, item := range slice {
		if len(item) < 3 || len(item) > 6 {
			t.Fatalf("StringSlice item length %d outside [3,6]", len(item))
		}
	}
}

func TestBytesZeroLength(t *testing.T) {
	b, err := Bytes(0)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if len(b) != 0 {
		t.Fatalf("Bytes length = %d want 0", len(b))
	}
	if _, err := Bytes(-1); err == nil {
		t.Fatalf("expected error for negative Bytes length")
	}
}
