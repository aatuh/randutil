package randtime

import (
	"math"
	"time"

	"github.com/aatuh/randutil/v2/core"
)

// Jitter returns base adjusted by a random factor in [-pct, +pct].
func Jitter(base time.Duration, pct float64) (time.Duration, error) {
	return Default().Jitter(base, pct)
}

// Jitter returns base adjusted by a random factor in [-pct, +pct].
func (g *Generator) Jitter(base time.Duration, pct float64) (time.Duration, error) {
	if base < 0 {
		return 0, core.ErrNegativeDuration
	}
	if pct < 0 || pct > 1 || math.IsNaN(pct) || math.IsInf(pct, 0) {
		return 0, core.ErrInvalidJitter
	}
	if pct == 0 || base == 0 {
		return base, nil
	}
	u, err := g.rng.Float64()
	if err != nil {
		return 0, err
	}
	delta := (u*2 - 1) * pct
	jittered := float64(base) * (1 + delta)
	return time.Duration(jittered), nil
}
