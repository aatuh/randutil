package dist

import (
	"errors"
	"math"
	"sync"

	"github.com/aatuh/randutil/v2/core"
)

var (
	errInvalidMeanStd      = errors.New("randutil: invalid mean/stddev")
	errNonFiniteParameter  = errors.New("randutil: parameter must be finite")
	errInvalidUniformRange = errors.New("randutil: min must be < max")
)

// Generator builds distribution samples using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng rng

	normalMu  sync.Mutex
	hasSpare  bool
	spareNorm float64
}

// New returns a dist Generator. If rng is nil, crypto/rand is used.
func New(rng rng) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator{rng: rng}
}

// NewWithSource returns a dist Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return New(core.New(src))
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// Bernoulli returns true with probability p and false otherwise using the generator's entropy source.
// p must be in [0,1].
func (g *Generator) Bernoulli(p float64) (bool, error) {
	if !isFinite(p) || p < 0 || p > 1 {
		return false, core.ErrInvalidProbability
	}
	u, err := g.rng.Float64()
	if err != nil {
		return false, err
	}
	return u < p, nil
}

// Categorical samples an index in [0, len(weights)) with probability
// proportional to weights[i] using the generator's entropy source.
// All weights must be >= 0 and at least one weight must be > 0.
func (g *Generator) Categorical(weights []float64) (int, error) {
	if len(weights) == 0 {
		return 0, core.ErrInvalidWeights
	}
	var sum float64
	for _, w := range weights {
		if !isFinite(w) || w < 0 {
			return 0, core.ErrInvalidWeights
		}
		sum += w
	}
	if sum <= 0 {
		return 0, core.ErrInvalidWeights
	}
	u, err := g.rng.Float64()
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
func (g *Generator) Exponential(lambda float64) (float64, error) {
	if !isFinite(lambda) {
		return 0, errNonFiniteParameter
	}
	if lambda <= 0 {
		return 0, core.ErrNonPositiveRate
	}
	u, err := g.rng.Float64()
	if err != nil {
		return 0, err
	}
	return -math.Log(1-u) / lambda, nil
}

// Normal returns a random value from a normal distribution
// with mean mu and standard deviation sigma using the generator's entropy source.
func (g *Generator) Normal(mu, sigma float64) (float64, error) {
	if !isFinite(mu) || !isFinite(sigma) {
		return 0, errInvalidMeanStd
	}
	if sigma < 0 {
		return 0, core.ErrNegativeStdDev
	}
	if sigma == 0 {
		return mu, nil
	}
	z, err := g.standardNormal()
	if err != nil {
		return 0, err
	}
	return mu + sigma*z, nil
}

func (g *Generator) standardNormal() (float64, error) {
	g.normalMu.Lock()
	if g.hasSpare {
		z := g.spareNorm
		g.hasSpare = false
		g.normalMu.Unlock()
		return z, nil
	}
	g.normalMu.Unlock()

	u1, err := g.rng.Float64()
	if err != nil {
		return 0, err
	}
	u2, err := g.rng.Float64()
	if err != nil {
		return 0, err
	}
	if u1 == 0 {
		u1 = math.SmallestNonzeroFloat64
	}
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	z0 := r * math.Cos(theta)
	z1 := r * math.Sin(theta)

	g.normalMu.Lock()
	g.spareNorm = z1
	g.hasSpare = true
	g.normalMu.Unlock()
	return z0, nil
}

// Uniform returns a random value from a uniform distribution
// in [min, max) using the generator's entropy source.
func (g *Generator) Uniform(minVal, maxVal float64) (float64, error) {
	if !isFinite(minVal) || !isFinite(maxVal) || minVal >= maxVal {
		return 0, errInvalidUniformRange
	}
	u, err := g.rng.Float64()
	if err != nil {
		return 0, err
	}
	return minVal + u*(maxVal-minVal), nil
}

// Poisson returns a random value from a Poisson distribution
// with parameter lambda using the generator's entropy source.
func (g *Generator) Poisson(lambda float64) (int, error) {
	if !isFinite(lambda) {
		return 0, errNonFiniteParameter
	}
	if lambda <= 0 {
		return 0, core.ErrNonPositiveRate
	}
	if lambda < 30 {
		return g.poissonKnuth(lambda)
	}
	return g.poissonPTRS(lambda)
}

func (g *Generator) poissonKnuth(lambda float64) (int, error) {
	L := math.Exp(-lambda)
	k := 0
	p := 1.0
	for p > L {
		u, err := g.rng.Float64()
		if err != nil {
			return 0, err
		}
		p *= u
		k++
	}
	return k - 1, nil
}

func (g *Generator) poissonPTRS(lambda float64) (int, error) {
	c := 0.767 - 3.36/lambda
	beta := math.Pi / math.Sqrt(3*lambda)
	alpha := beta * lambda
	k := math.Log(c) - lambda - math.Log(beta)
	for {
		u, err := g.rng.Float64()
		if err != nil {
			return 0, err
		}
		if u <= 0 || u >= 1 {
			continue
		}
		x := (alpha - math.Log((1-u)/u)) / beta
		n := math.Floor(x + 0.5)
		if n < 0 {
			continue
		}
		v, err := g.rng.Float64()
		if err != nil {
			return 0, err
		}
		if v <= 0 {
			continue
		}
		y := alpha - beta*x
		lhs := y + math.Log(v) - 2*log1pexp(y)
		// Lgamma sign is always +1 for positive integer inputs.
		lg, _ := math.Lgamma(n + 1)
		rhs := k + n*math.Log(lambda) - lg
		if lhs <= rhs {
			return int(n), nil
		}
	}
}

// Gamma returns a random value from a gamma distribution
// with shape alpha and rate beta using the generator's entropy source.
func (g *Generator) Gamma(alpha, beta float64) (float64, error) {
	if !isFinite(alpha) || !isFinite(beta) {
		return 0, errNonFiniteParameter
	}
	if alpha <= 0 {
		return 0, core.ErrNonPositiveBound
	}
	if beta <= 0 {
		return 0, core.ErrNonPositiveRate
	}
	x, err := g.gammaStandard(alpha)
	if err != nil {
		return 0, err
	}
	return x / beta, nil
}

func (g *Generator) gammaStandard(alpha float64) (float64, error) {
	if !isFinite(alpha) {
		return 0, errNonFiniteParameter
	}
	if alpha <= 0 {
		return 0, core.ErrNonPositiveBound
	}
	if alpha < 1 {
		u, err := g.rng.Float64()
		if err != nil {
			return 0, err
		}
		if u == 0 {
			u = math.SmallestNonzeroFloat64
		}
		x, err := g.gammaStandard(alpha + 1)
		if err != nil {
			return 0, err
		}
		return x * math.Pow(u, 1/alpha), nil
	}

	d := alpha - 1.0/3.0
	c := 1 / math.Sqrt(9*d)
	for {
		x, err := g.standardNormal()
		if err != nil {
			return 0, err
		}
		v := 1 + c*x
		if v <= 0 {
			continue
		}
		v = v * v * v
		u, err := g.rng.Float64()
		if err != nil {
			return 0, err
		}
		if u < 1-0.0331*(x*x)*(x*x) {
			return d * v, nil
		}
		if math.Log(u) < 0.5*x*x+d*(1-v+math.Log(v)) {
			return d * v, nil
		}
	}
}

func log1pexp(x float64) float64 {
	if x > 0 {
		return x + math.Log1p(math.Exp(-x))
	}
	return math.Log1p(math.Exp(x))
}

func isFinite(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0)
}
