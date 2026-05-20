package adapters

import (
	"sync"

	"golang.org/x/crypto/chacha20"

	"github.com/aatuh/randutil/v2/core"
)

type chachaSource struct {
	mu     sync.Mutex
	cipher *chacha20.Cipher
	key    [32]byte
	nonce  [12]byte
	closed bool
}

func newChaChaSource(key [32]byte, nonce [12]byte) (core.Source, error) {
	cipher, err := chacha20.NewUnauthenticatedCipher(key[:], nonce[:])
	if err != nil {
		return nil, err
	}
	return &chachaSource{
		cipher: cipher,
		key:    key,
		nonce:  nonce,
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
	for i := range p {
		p[i] = 0
	}
	c.cipher.XORKeyStream(p, p)
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
	core.Zero(c.key[:])
	core.Zero(c.nonce[:])
	return nil
}
