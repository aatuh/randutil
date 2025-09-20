package core

import (
	"crypto/rand"
	"io"
	"sync"
)

// sourceMux guards currentSource to allow swapping entropy sources
// (e.g., deterministic readers in tests). By default, crypto/rand is
// used. All public funcs should read via getSource().
var (
	sourceMux     sync.RWMutex
	currentSource io.Reader = rand.Reader
)

// SetSource changes the active entropy source. It is safe for
// concurrent use. Do not pass nil.
//
// Parameters:
//   - r: The new entropy source.
func SetSource(r io.Reader) {
	if r == nil {
		return
	}
	sourceMux.Lock()
	currentSource = r
	sourceMux.Unlock()
}

// GetSource is for fetching the current source quickly.
//
// Returns:
//   - io.Reader: The current source.
func GetSource() io.Reader {
	sourceMux.RLock()
	r := currentSource
	sourceMux.RUnlock()
	return r
}

// ResetSource restores the default crypto/rand.Reader.
//
// Parameters:
//   - r: The new entropy source.
func ResetSource() {
	sourceMux.Lock()
	currentSource = rand.Reader
	sourceMux.Unlock()
}

// Reader returns an io.Reader proxy that delegates to the current
// source. Useful when integrating with code that accepts a reader for
// entropy (e.g., TLS, keygens, or external libs).
//
// Returns:
//   - io.Reader: The current source.
func Reader() io.Reader {
	return readerProxy{}
}

// readerProxy is a proxy for the current source.
type readerProxy struct{}

// Read reads from the current source.
//
// Parameters:
//   - p: A byte slice to read into.
//
// Returns:
//   - int: The number of bytes read.
//   - error: An error if the source fails.
func (readerProxy) Read(p []byte) (int, error) {
	sourceMux.RLock()
	r := currentSource
	sourceMux.RUnlock()
	return r.Read(p)
}
