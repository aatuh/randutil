package ulid

import (
	"encoding/binary"
	"time"

	"github.com/aatuh/randutil/v2/core"
)

const maxULIDTime = int64(1<<48 - 1)

// Generator builds ULIDs using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng rng
	now func() time.Time
}

// New returns a ULID Generator. If rng is nil, crypto/rand is used.
func New(rng rng) *Generator {
	return NewWithClock(rng, time.Now)
}

// NewWithClock returns a ULID Generator bound to rng and clock.
// If rng is nil, crypto/rand is used. If now is nil, time.Now is used.
func NewWithClock(rng rng, now func() time.Time) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	if now == nil {
		now = time.Now
	}
	return &Generator{rng: rng, now: now}
}

// NewWithSource returns a ULID Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return NewWithClock(core.New(src), time.Now)
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// ULID returns a ULID based on the current time and random data.
func (g *Generator) ULID() (ULID, error) {
	ms := g.nowUTC().UnixMilli()
	if ms < 0 || ms > maxULIDTime {
		return "", core.ErrResultOutOfRange
	}
	var buf [16]byte
	var ts [8]byte
	binary.BigEndian.PutUint64(ts[:], uint64(ms))
	copy(buf[:6], ts[2:])
	if err := g.rng.Fill(buf[6:]); err != nil {
		return "", err
	}
	return ULID(ulidEncoding.EncodeToString(buf[:])), nil
}

func (g *Generator) nowUTC() time.Time {
	if g == nil || g.now == nil {
		return time.Now().UTC()
	}
	return g.now().UTC()
}
