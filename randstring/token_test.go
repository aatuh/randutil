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

	hb, err := TokenHexBytes(sz)
	if err != nil {
		t.Fatalf("TokenHexBytes error: %v", err)
	}
	if len(hb) != 2*sz {
		t.Fatalf("hex bytes length = %d want %d", len(hb), 2*sz)
	}
	if _, err := hex.DecodeString(string(hb)); err != nil {
		t.Fatalf("invalid hex bytes: %v", err)
	}

	b64, err := TokenBase64(sz)
	if err != nil {
		t.Fatalf("TokenBase64 error: %v", err)
	}
	if _, err := base64.StdEncoding.DecodeString(b64); err != nil {
		t.Fatalf("invalid base64: %v", err)
	}

	b64b, err := TokenBase64Bytes(sz)
	if err != nil {
		t.Fatalf("TokenBase64Bytes error: %v", err)
	}
	if _, err := base64.StdEncoding.DecodeString(string(b64b)); err != nil {
		t.Fatalf("invalid base64 bytes: %v", err)
	}

	url, err := TokenURLSafe(sz)
	if err != nil {
		t.Fatalf("TokenURLSafe error: %v", err)
	}
	if _, err := base64.RawURLEncoding.DecodeString(url); err != nil {
		t.Fatalf("invalid url-safe b64: %v", err)
	}

	urlb, err := TokenURLSafeBytes(sz)
	if err != nil {
		t.Fatalf("TokenURLSafeBytes error: %v", err)
	}
	if _, err := base64.RawURLEncoding.DecodeString(string(urlb)); err != nil {
		t.Fatalf("invalid url-safe b64 bytes: %v", err)
	}
}
