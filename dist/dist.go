package dist

import (
	"errors"
	"math"
	"time"
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
	return Default.Bernoulli(p)
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
	return Default.Categorical(weights)
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
	return Default.Exponential(lambda)
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
	return Default.Normal(mean, stddev)
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
	u, err := Default.G.Float64()
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

// MustBernoulli returns true with probability p and false otherwise.
// It panics on error.
//
// Parameters:
//   - p: The probability of true.
//
// Returns:
//   - bool: A random boolean value.
func MustBernoulli(p float64) bool {
	return Default.MustBernoulli(p)
}

// MustCategorical samples an index in [0, len(weights)) with probability
// proportional to weights[i]. It panics on error.
//
// Parameters:
//   - weights: The weights of the categories.
//
// Returns:
//   - int: A random index in [0, len(weights)).
func MustCategorical(weights []float64) int {
	return Default.MustCategorical(weights)
}

// MustExponential returns a random value from an exponential distribution
// with rate parameter lambda. It panics on error.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the exponential distribution.
func MustExponential(lambda float64) float64 {
	return Default.MustExponential(lambda)
}

// MustNormal returns a random value from a normal distribution
// with mean mu and standard deviation sigma. It panics on error.
//
// Parameters:
//   - mu: The mean of the distribution.
//   - sigma: The standard deviation (must be > 0).
//
// Returns:
//   - float64: A random value from the normal distribution.
func MustNormal(mu, sigma float64) float64 {
	return Default.MustNormal(mu, sigma)
}

// MustUniform returns a random value from a uniform distribution
// in [min, max). It panics on error.
//
// Parameters:
//   - min: The minimum value (inclusive).
//   - max: The maximum value (exclusive).
//
// Returns:
//   - float64: A random value from the uniform distribution.
func MustUniform(min, max float64) float64 {
	return Default.MustUniform(min, max)
}

// MustPoisson returns a random value from a Poisson distribution
// with parameter lambda. It panics on error.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - int: A random value from the Poisson distribution.
func MustPoisson(lambda float64) int {
	return Default.MustPoisson(lambda)
}

// MustGamma returns a random value from a gamma distribution
// with shape alpha and rate beta. It panics on error.
//
// Parameters:
//   - alpha: The shape parameter (must be > 0).
//   - beta: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the gamma distribution.
func MustGamma(alpha, beta float64) float64 {
	return Default.MustGamma(alpha, beta)
}
