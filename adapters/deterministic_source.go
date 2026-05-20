//go:build !randutil_policy
// +build !randutil_policy

package adapters

import (
	"crypto/sha256"

	"github.com/aatuh/randutil/v2/core"
)

// DeterministicSource returns a reproducible stream based on seed.
//
// WARNING: This is deterministic. DO NOT USE FOR TOKENS / AUTH unless the seed
// is high-entropy and kept secret. Intended for tests, benchmarks, and
// replayable simulations.
//
// Returns an error when policy mode disables deterministic sources.
func DeterministicSource(seed []byte) (core.Source, error) {
	return DeterministicSourceWithLabel(seed, "")
}

// DeterministicSourceWithLabel returns a reproducible stream based on seed and label.
// Use labels to derive independent deterministic streams from the same seed.
//
// WARNING: This is deterministic. DO NOT USE FOR TOKENS / AUTH unless the seed
// is high-entropy and kept secret. Intended for tests, benchmarks, and
// replayable simulations.
//
// Returns an error when policy mode disables deterministic sources.
func DeterministicSourceWithLabel(seed []byte, label string) (core.Source, error) {
	if label == "" {
		return deterministicSourceFromSeed(seed)
	}
	derived := deriveLabelSeed(seed, label)
	src, err := deterministicSourceFromSeed(derived[:])
	core.Zero(derived[:])
	return src, err
}

func deterministicSourceFromSeed(seed []byte) (core.Source, error) {
	key := deriveKey(seed, "key")
	nonceKey := deriveKey(seed, "nonce")
	var nonce [12]byte
	copy(nonce[:], nonceKey[:12])
	src, err := newChaChaSource(key, nonce)
	core.Zero(nonceKey[:])
	return src, err
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
