package dist

import (
	"errors"
	"math"
	"time"

	"github.com/aatuh/randutil/numeric"
)

// Bernoulli returns true with probability p and false otherwise.
// p must be in [0,1].
//
// Parameters:
//   - p: The probability of true.
//
// Returns:
//   - bool: A random boolean value.
//   - error: An error if the source fails.
func Bernoulli(p float64) (bool, error) {
	if math.IsNaN(p) || p < 0 || p > 1 {
		return false, errors.New("p must be in [0,1]")
	}
	u, err := numeric.Float64()
	if err != nil {
		return false, err
	}
	return u < p, nil
}

// Categorical samples an index in [0, len(weights)) with probability
// proportional to weights[i]. All weights must be >= 0 and at least
// one weight must be > 0.
//
// Parameters:
//   - weights: The weights of the categories.
//
// Returns:
//   - int: A random index in [0, len(weights)).
//   - error: An error if the source fails.
func Categorical(weights []float64) (int, error) {
	var sum float64
	for _, w := range weights {
		if math.IsNaN(w) || w < 0 {
			return 0, errors.New("weights must be >= 0")
		}
		sum += w
	}
	if sum == 0 {
		return 0, errors.New("at least one weight must be > 0")
	}
	u, err := numeric.Float64()
	if err != nil {
		return 0, err
	}
	target := u * sum
	acc := 0.0
	for i, w := range weights {
		acc += w
		if target < acc {
			return i, nil
		}
	}
	return len(weights) - 1, nil
}

// Exponential samples an exponential(lambda). lambda must be > 0.
//
// Parameters:
//   - lambda: The lambda of the exponential distribution.
//
// Returns:
//   - float64: A random float64 value.
//   - error: An error if the source fails.
func Exponential(lambda float64) (float64, error) {
	if !(lambda > 0) || math.IsNaN(lambda) {
		return 0, errors.New("lambda must be > 0")
	}
	u, err := numeric.Float64()
	if err != nil {
		return 0, err
	}
	// Use 1-u to avoid log(0); u in [0,1).
	if u == 0 {
		u = math.SmallestNonzeroFloat64
	}
	return -math.Log(1-u) / lambda, nil
}

// Normal returns a normal(mean, stddev) variate using Box-Muller.
// stddev must be >= 0. If stddev == 0, returns mean.
//
// Parameters:
//   - mean: The mean of the normal distribution.
//   - stddev: The standard deviation of the normal distribution.
//
// Returns:
//   - float64: A random float64 value.
//   - error: An error if the source fails.
func Normal(mean, stddev float64) (float64, error) {
	if math.IsNaN(mean) || math.IsNaN(stddev) || stddev < 0 {
		return 0, errors.New("invalid mean/stddev")
	}
	if stddev == 0 {
		return mean, nil
	}
	u1, err := numeric.Float64()
	if err != nil {
		return 0, err
	}
	u2, err := numeric.Float64()
	if err != nil {
		return 0, err
	}
	if u1 == 0 {
		u1 = math.SmallestNonzeroFloat64
	}
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	z := r * math.Cos(theta)
	return mean + stddev*z, nil
}

// Zipf is a precomputed sampler for Zipf(s, v) over [1..imax] where:
//
//	P(X=k) ∝ (v + k)^(-s).
//
// It builds a normalized CDF for O(log n) sampling via binary search.
type Zipf struct {
	s, v  float64
	imax  int
	cdf   []float64
	total float64
}

// NewZipf builds a Zipf sampler. s > 1 is typical; v >= 0; imax >= 1.
//
// Parameters:
//   - s: The s of the Zipf distribution.
//   - v: The v of the Zipf distribution.
//   - imax: The imax of the Zipf distribution.
//
// Returns:
//   - *Zipf: A new Zipf sampler.
//   - error: An error if the source fails.
func NewZipf(s, v float64, imax int) (*Zipf, error) {
	if math.IsNaN(s) || math.IsNaN(v) || s <= 0 || v < 0 || imax < 1 {
		return nil, errors.New("invalid s, v, or imax")
	}
	z := &Zipf{s: s, v: v, imax: imax}
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
//
// Returns:
//   - int: A random index in [1..imax].
//   - error: An error if the source fails.
func (z *Zipf) Next() (int, error) {
	u, err := numeric.Float64()
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

// SeededClockNormal returns a normal variate around time.Now with
// jitter stddev seconds. It is a small example of composing dists.
//
// Parameters:
//   - stddevSeconds: The standard deviation of the normal distribution.
//
// Returns:
//   - time.Time: A random time.Time value.
//   - error: An error if the source fails.
func SeededClockNormal(stddevSeconds float64) (time.Time, error) {
	now := time.Now().UTC()
	j, err := Normal(0, stddevSeconds)
	if err != nil {
		return time.Time{}, err
	}
	return now.Add(time.Duration(j * float64(time.Second))), nil
}
