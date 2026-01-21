package adapters

import (
	"crypto/sha256"
	"sync"

	"golang.org/x/crypto/chacha20"

	"github.com/aatuh/randutil/v2/core"
)

type deterministicSource struct {
	mu     sync.Mutex
	cipher *chacha20.Cipher
}

// DeterministicSource returns a reproducible stream based on seed.
//
// WARNING: This is deterministic. Do not use for security unless the seed is
// high-entropy and kept secret. Intended for tests, benchmarks, and replayable
// simulations.
func DeterministicSource(seed []byte) core.Source {
	return DeterministicSourceWithLabel(seed, "")
}

// DeterministicSourceWithLabel returns a reproducible stream based on seed and label.
// Use labels to derive independent deterministic streams from the same seed.
//
// WARNING: This is deterministic. Do not use for security unless the seed is
// high-entropy and kept secret. Intended for tests, benchmarks, and replayable
// simulations.
func DeterministicSourceWithLabel(seed []byte, label string) core.Source {
	if label == "" {
		return deterministicSourceFromSeed(seed)
	}
	derived := deriveLabelSeed(seed, label)
	return deterministicSourceFromSeed(derived[:])
}

func deterministicSourceFromSeed(seed []byte) core.Source {
	key := deriveKey(seed, "key")
	nonce := deriveKey(seed, "nonce")
	cipher, err := chacha20.NewUnauthenticatedCipher(key[:], nonce[:12])
	if err != nil {
		panic("randutil: deterministic source cipher init failed: " + err.Error())
	}
	return &deterministicSource{cipher: cipher}
}

func (d *deterministicSource) Read(p []byte) (int, error) {
	d.mu.Lock()
	defer d.mu.Unlock()
	for i := range p {
		p[i] = 0
	}
	d.cipher.XORKeyStream(p, p)
	return len(p), nil
}

func deriveKey(seed []byte, label string) [32]byte {
	h := sha256.New()
	h.Write([]byte("randutil deterministic source "))
	h.Write([]byte(label))
	h.Write(seed)
	var out [32]byte
	sum := h.Sum(nil)
	copy(out[:], sum)
	return out
}

func deriveLabelSeed(seed []byte, label string) [32]byte {
	h := sha256.New()
	h.Write([]byte("randutil deterministic source label "))
	h.Write([]byte(label))
	h.Write(seed)
	var out [32]byte
	sum := h.Sum(nil)
	copy(out[:], sum)
	return out
}
