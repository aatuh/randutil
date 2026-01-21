package adapters

import (
	"crypto/rand"

	"github.com/aatuh/randutil/v2/core"
)

// CryptoSource returns a Source backed by crypto/rand.Reader.
func CryptoSource() core.Source {
	return rand.Reader
}
