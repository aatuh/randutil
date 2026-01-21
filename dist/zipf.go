package dist

import (
	"errors"
	"math"

	"github.com/aatuh/randutil/v2/core"
)

// Zipf is a precomputed sampler for Zipf(s, v) over [1..imax] where:
//
//	P(X=k) ∝ (v + k)^(-s).
//
// It builds a normalized CDF for O(log n) sampling via binary search.
type Zipf struct {
	rng   core.RNG
	s     float64
	v     float64
	imax  int
	cdf   []float64
	total float64
}

// NewZipf builds a Zipf sampler using the default generator. s > 1 is typical;
// v >= 0; imax >= 1.
func NewZipf(s, v float64, imax int) (*Zipf, error) {
	return Default().Zipf(s, v, imax)
}

// Zipf builds a Zipf sampler using the generator's entropy source.
func (g *Generator) Zipf(s, v float64, imax int) (*Zipf, error) {
	if math.IsNaN(s) || math.IsNaN(v) || s <= 0 || v < 0 || imax < 1 {
		return nil, errors.New("randutil: invalid s, v, or imax")
	}
	z := &Zipf{rng: g.rng, s: s, v: v, imax: imax}
	z.cdf = make([]float64, imax)
	var acc float64
	for k := 1; k <= imax; k++ {
		acc += math.Pow(z.v+float64(k), -z.s)
		z.cdf[k-1] = acc
	}
	z.total = acc
	for i := range z.cdf {
		z.cdf[i] /= z.total
	}
	return z, nil
}

// Next draws one sample in [1..imax].
func (z *Zipf) Next() (int, error) {
	if z == nil || z.rng == nil {
		return 0, errors.New("randutil: nil Zipf rng")
	}
	u, err := z.rng.Float64()
	if err != nil {
		return 0, err
	}
	lo, hi := 0, len(z.cdf)-1
	for lo < hi {
		mid := (lo + hi) / 2
		if u <= z.cdf[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	return lo + 1, nil
}
