package uuid

import (
	"time"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds UUID-related random operations using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng core.RNG
	now func() time.Time
}

// New returns a uuid Generator. If rng is nil, crypto/rand is used.
func New(rng core.RNG) *Generator {
	return NewWithClock(rng, time.Now)
}

// NewWithClock returns a uuid Generator bound to rng and clock.
// If rng is nil, crypto/rand is used. If now is nil, time.Now is used.
func NewWithClock(rng core.RNG, now func() time.Time) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	if now == nil {
		now = time.Now
	}
	return &Generator{rng: rng, now: now}
}

// NewWithSource returns a uuid Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return NewWithClock(core.New(src), time.Now)
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// V4 returns a RFC 4122, variant 1 UUID v4 using the generator's entropy source.
// It reads 16 bytes, then sets version and variant bits.
//
// Returns:
//   - UUID: A random UUID conforming to Version 4 and Variant 1.
//   - error: An error if entropy fails.
func (g *Generator) V4() (UUID, error) {
	b, err := g.rng.Bytes(16)
	if err != nil {
		return "", err
	}
	b[6] = (b[6] & 0x0f) | 0x40 // version 4
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10xx
	return fromBytes(b), nil
}

// MustV4 returns a v4 UUID or panics on error.
func (g *Generator) MustV4() UUID {
	u, err := g.V4()
	if err != nil {
		panic(err)
	}
	return u
}

// V7 returns a RFC 9562, variant 1 UUID v7 (time-ordered) using the generator's entropy source.
// The first 48 bits encode Unix milliseconds big-endian.
//
// Returns:
//   - UUID: A random UUID conforming to Version 7 and Variant 1.
//   - error: An error if entropy fails.
func (g *Generator) V7() (UUID, error) {
	b, err := g.rng.Bytes(16)
	if err != nil {
		return "", err
	}
	ms := g.nowUTC().UnixMilli()
	if ms < 0 {
		return "", core.ErrResultOutOfRange
	}
	msu := uint64(ms)
	b[0] = byte(msu >> 40)
	b[1] = byte(msu >> 32)
	b[2] = byte(msu >> 24)
	b[3] = byte(msu >> 16)
	b[4] = byte(msu >> 8)
	b[5] = byte(msu)
	b[6] = (b[6] & 0x0f) | 0x70 // version 7
	b[8] = (b[8] & 0x3f) | 0x80 // variant 10xx
	return fromBytes(b), nil
}

// MustV7 returns a v7 UUID or panics on error.
func (g *Generator) MustV7() UUID {
	u, err := g.V7()
	if err != nil {
		panic(err)
	}
	return u
}

func (g *Generator) nowUTC() time.Time {
	if g == nil || g.now == nil {
		return time.Now().UTC()
	}
	return g.now().UTC()
}
