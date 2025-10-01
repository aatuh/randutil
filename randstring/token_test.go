package randstring

import (
	"encoding/base64"
	"encoding/hex"
	"testing"
)

func TestTokens(t *testing.T) {
	const sz = 24
	h, err := TokenHex(sz)
	if err != nil {
		t.Fatalf("TokenHex error: %v", err)
	}
	// Hex length is 2*n.
	if len(h) != 2*sz {
		t.Fatalf("hex length = %d want %d", len(h), 2*sz)
	}
	if _, err := hex.DecodeString(h); err != nil {
		t.Fatalf("invalid hex: %v", err)
	}

	b64, err := TokenBase64(sz)
	if err != nil {
		t.Fatalf("TokenBase64 error: %v", err)
	}
	if _, err := base64.StdEncoding.DecodeString(b64); err != nil {
		t.Fatalf("invalid base64: %v", err)
	}

	url, err := TokenURLSafe(sz)
	if err != nil {
		t.Fatalf("TokenURLSafe error: %v", err)
	}
	if _, err := base64.RawURLEncoding.DecodeString(url); err != nil {
		t.Fatalf("invalid url-safe b64: %v", err)
	}
}
