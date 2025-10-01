package dist

import (
	"errors"
	"io"
	"math"

	"github.com/aatuh/randutil/v2/core"
)

// Generator builds distribution samples using a core generator.
type Generator struct {
	G core.Generator
}

// New returns a dist Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator { return &Generator{G: core.Generator{R: src}} }

// Default is the package-wide default generator.
var Default = New(nil)

// Bernoulli returns true with probability p and false otherwise using the generator's entropy source.
// p must be in [0,1].
//
// Parameters:
//   - p: The probability of true.
//
// Returns:
//   - bool: A random boolean value.
//   - error: An error if the source fails.
func (g *Generator) Bernoulli(p float64) (bool, error) {
	if math.IsNaN(p) || p < 0 || p > 1 {
		return false, errors.New("p must be in [0,1]")
	}
	u, err := g.G.Float64()
	if err != nil {
		return false, err
	}
	return u < p, nil
}

// Categorical samples an index in [0, len(weights)) with probability
// proportional to weights[i] using the generator's entropy source.
// All weights must be >= 0 and at least one weight must be > 0.
//
// Parameters:
//   - weights: The weights of the categories.
//
// Returns:
//   - int: A random index in [0, len(weights)).
//   - error: An error if the source fails.
func (g *Generator) Categorical(weights []float64) (int, error) {
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
	u, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	target := u * sum
	var acc float64
	for i, w := range weights {
		acc += w
		if target < acc {
			return i, nil
		}
	}
	return len(weights) - 1, nil
}

// Exponential returns a random value from an exponential distribution
// with rate parameter lambda using the generator's entropy source.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the exponential distribution.
//   - error: An error if the source fails.
func (g *Generator) Exponential(lambda float64) (float64, error) {
	if lambda <= 0 {
		return 0, errors.New("lambda must be > 0")
	}
	u, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	return -math.Log(1-u) / lambda, nil
}

// Normal returns a random value from a normal distribution
// with mean mu and standard deviation sigma using the generator's entropy source.
//
// Parameters:
//   - mu: The mean of the distribution.
//   - sigma: The standard deviation (must be >= 0).
//
// Returns:
//   - float64: A random value from the normal distribution.
//   - error: An error if the source fails.
func (g *Generator) Normal(mu, sigma float64) (float64, error) {
	if math.IsNaN(mu) || math.IsNaN(sigma) || sigma < 0 {
		return 0, errors.New("invalid mean/stddev")
	}
	if sigma == 0 {
		return mu, nil
	}
	u1, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	u2, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	if u1 == 0 {
		u1 = math.SmallestNonzeroFloat64
	}
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	z := r * math.Cos(theta)
	return mu + sigma*z, nil
}

// Uniform returns a random value from a uniform distribution
// in [min, max) using the generator's entropy source.
//
// Parameters:
//   - min: The minimum value (inclusive).
//   - max: The maximum value (exclusive).
//
// Returns:
//   - float64: A random value from the uniform distribution.
//   - error: An error if the source fails.
func (g *Generator) Uniform(min, max float64) (float64, error) {
	if min >= max {
		return 0, errors.New("min must be < max")
	}
	u, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	return min + u*(max-min), nil
}

// Poisson returns a random value from a Poisson distribution
// with parameter lambda using the generator's entropy source.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - int: A random value from the Poisson distribution.
//   - error: An error if the source fails.
func (g *Generator) Poisson(lambda float64) (int, error) {
	if lambda <= 0 {
		return 0, errors.New("lambda must be > 0")
	}
	// Knuth's algorithm
	L := math.Exp(-lambda)
	k := 0
	p := 1.0
	for p > L {
		u, err := g.G.Float64()
		if err != nil {
			return 0, err
		}
		p *= u
		k++
	}
	return k - 1, nil
}

// Gamma returns a random value from a gamma distribution
// with shape alpha and rate beta using the generator's entropy source.
//
// Parameters:
//   - alpha: The shape parameter (must be > 0).
//   - beta: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the gamma distribution.
//   - error: An error if the source fails.
func (g *Generator) Gamma(alpha, beta float64) (float64, error) {
	if alpha <= 0 || beta <= 0 {
		return 0, errors.New("alpha and beta must be > 0")
	}
	// Simple rejection sampling for alpha >= 1
	if alpha >= 1 {
		return g.gammaRejection(alpha, beta)
	}
	// For alpha < 1, use the relationship Gamma(alpha, beta) = Gamma(alpha+1, beta) * U^(1/alpha)
	// where U is uniform(0,1)
	u, err := g.G.Float64()
	if err != nil {
		return 0, err
	}
	g1, err := g.gammaRejection(alpha+1, beta)
	if err != nil {
		return 0, err
	}
	return g1 * math.Pow(u, 1.0/alpha), nil
}

// gammaRejection implements rejection sampling for gamma distribution.
func (g *Generator) gammaRejection(alpha, beta float64) (float64, error) {
	// This is a simplified implementation
	// In practice, you'd want a more sophisticated algorithm
	for {
		u, err := g.G.Float64()
		if err != nil {
			return 0, err
		}
		v, err := g.G.Float64()
		if err != nil {
			return 0, err
		}
		x := -math.Log(u)
		if v <= math.Pow(x, alpha-1) {
			return x / beta, nil
		}
	}
}

// MustBernoulli returns true with probability p and false otherwise using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - p: The probability of true.
//
// Returns:
//   - bool: A random boolean value.
func (g *Generator) MustBernoulli(p float64) bool {
	b, err := g.Bernoulli(p)
	if err != nil {
		panic(err)
	}
	return b
}

// MustCategorical samples an index in [0, len(weights)) with probability
// proportional to weights[i] using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - weights: The weights of the categories.
//
// Returns:
//   - int: A random index in [0, len(weights)).
func (g *Generator) MustCategorical(weights []float64) int {
	i, err := g.Categorical(weights)
	if err != nil {
		panic(err)
	}
	return i
}

// MustExponential returns a random value from an exponential distribution
// with rate parameter lambda using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the exponential distribution.
func (g *Generator) MustExponential(lambda float64) float64 {
	f, err := g.Exponential(lambda)
	if err != nil {
		panic(err)
	}
	return f
}

// MustNormal returns a random value from a normal distribution
// with mean mu and standard deviation sigma using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - mu: The mean of the distribution.
//   - sigma: The standard deviation (must be > 0).
//
// Returns:
//   - float64: A random value from the normal distribution.
func (g *Generator) MustNormal(mu, sigma float64) float64 {
	f, err := g.Normal(mu, sigma)
	if err != nil {
		panic(err)
	}
	return f
}

// MustUniform returns a random value from a uniform distribution
// in [min, max) using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - min: The minimum value (inclusive).
//   - max: The maximum value (exclusive).
//
// Returns:
//   - float64: A random value from the uniform distribution.
func (g *Generator) MustUniform(min, max float64) float64 {
	f, err := g.Uniform(min, max)
	if err != nil {
		panic(err)
	}
	return f
}

// MustPoisson returns a random value from a Poisson distribution
// with parameter lambda using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - lambda: The rate parameter (must be > 0).
//
// Returns:
//   - int: A random value from the Poisson distribution.
func (g *Generator) MustPoisson(lambda float64) int {
	i, err := g.Poisson(lambda)
	if err != nil {
		panic(err)
	}
	return i
}

// MustGamma returns a random value from a gamma distribution
// with shape alpha and rate beta using the generator's entropy source.
// It panics on error.
//
// Parameters:
//   - alpha: The shape parameter (must be > 0).
//   - beta: The rate parameter (must be > 0).
//
// Returns:
//   - float64: A random value from the gamma distribution.
func (g *Generator) MustGamma(alpha, beta float64) float64 {
	f, err := g.Gamma(alpha, beta)
	if err != nil {
		panic(err)
	}
	return f
}
