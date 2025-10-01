package uuid

import (
	"io"
	"time"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds UUID-related random operations using a core generator.
type Generator struct {
	G core.Generator
}

// New returns a uuid Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator {
	return &Generator{G: core.Generator{R: src}}
}

// Default is the package-wide default generator.
var Default = New(nil)

// V4 returns a RFC 4122, variant 1 UUID v4 using the generator's entropy source.
// It reads 16 bytes, then sets version and variant bits.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
//   - error: An error if entropy fails.
func (g *Generator) V4() (UUID, error) {
	b, err := g.G.Bytes(16)
	if err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10xx
	return fromBytes(b), nil
}

// V7 returns a RFC 9562, variant 1 UUID v7 (time-ordered) using the generator's entropy source.
// The first 48 bits encode Unix milliseconds big-endian.
//
// Returns:
//   - UUID: A random UUID conforming to Version 7 and Variant 1.
//   - error: An error if entropy fails.
func (g *Generator) V7() (UUID, error) {
	b, err := g.G.Bytes(16)
	if err != nil {
		return "", err
	}
	ms := uint64(time.Now().UTC().UnixMilli())
	b[0] = byte(ms >> 40)
	b[1] = byte(ms >> 32)
	b[2] = byte(ms >> 24)
	b[3] = byte(ms >> 16)
	b[4] = byte(ms >> 8)
	b[5] = byte(ms)
	b[6] = (b[6] & 0x0f) | 0x70 // version 7
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10xx
	return fromBytes(b), nil
}
