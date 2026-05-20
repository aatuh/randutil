package adapters

import (
	"sync"

	"golang.org/x/crypto/chacha20"

	"github.com/aatuh/randutil/v2/core"
)

const maxChaChaSourceBytes = uint64(1<<32) * 64

type chachaSource struct {
	mu     sync.Mutex
	cipher *chacha20.Cipher
	key    [32]byte
	nonce  [12]byte
	used   uint64
	limit  uint64
	closed bool
}

func newChaChaSource(key [32]byte, nonce [12]byte) (core.Source, error) {
	return newChaChaSourceWithLimit(key, nonce, maxChaChaSourceBytes)
}

func newChaChaSourceWithLimit(key [32]byte, nonce [12]byte, limit uint64) (core.Source, error) {
	cipher, err := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	if err != nil {
		return nil, err
	}
	return &chachaSource{
		cipher: cipher,
		key:    key,
		nonce:  nonce,
		limit:  limit,
	}, nil
}

func (c *chachaSource) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed || c.cipher == nil {
		return 0, core.ErrSourceClosed
	}
	if uint64(len(p)) > c.limit-c.used {
		for i := range p {
			p[i] = 0
		}
		return 0, core.ErrSourceExhausted
	}
	for i := range p {
		p[i] = 0
	}
	c.cipher.XORKeyStream(p, p)
	c.used += uint64(len(p))
	return len(p), nil
}

func (c *chachaSource) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.closed {
		return nil
	}
	c.closed = true
	c.cipher = nil
	c.used = c.limit
	core.Zero(c.key[:])
	core.Zero(c.nonce[:])
	return nil
}
