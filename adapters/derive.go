package adapters

import (
	"crypto/sha256"
	"io"

	"golang.org/x/crypto/hkdf"

	"github.com/aatuh/randutil/v2/core"
)

const (
	deriveSalt       = "randutil hkdf v1"
	deriveInfoPrefix = "randutil derive v1 "
)

// DeriveSource returns a domain-separated stream derived from seed and label.
// The stream is cryptographically strong if the seed is high entropy and secret.
// It uses HKDF-SHA256 + ChaCha20; for FIPS/OS RNG compliance use crypto/rand
// directly instead of derived streams.
func DeriveSource(seed []byte, label string) (core.Source, error) {
	key, nonce, err := deriveKeyNonce(seed, label)
	if err != nil {
		return nil, err
	}
	return newChaChaSource(key, nonce)
}

// DeriveRNG returns a core RNG derived from seed and label.
func DeriveRNG(seed []byte, label string) (core.RNG, error) {
	src, err := DeriveSource(seed, label)
	if err != nil {
		return nil, err
	}
	return core.New(src), nil
}

func deriveKeyNonce(seed []byte, label string) ([32]byte, [12]byte, error) {
	info := []byte(deriveInfoPrefix)
	if label != "" {
		info = append(info, label...)
	}
	reader := hkdf.New(sha256.New, seed, []byte(deriveSalt), info)
	var out [44]byte
	if _, err := io.ReadFull(reader, out[:]); err != nil {
		return [32]byte{}, [12]byte{}, err
	}
	var key [32]byte
	var nonce [12]byte
	copy(key[:], out[:32])
	copy(nonce[:], out[32:])
	core.Zero(out[:])
	return key, nonce, nil
}
