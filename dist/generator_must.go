//go:build randutil_must
// +build randutil_must

package dist

// MustBernoulli returns true with probability p and false otherwise using the generator's entropy source.
// It panics on error.
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
func (g *Generator) MustUniform(minVal, maxVal float64) float64 {
	f, err := g.Uniform(minVal, maxVal)
	if err != nil {
		panic(err)
	}
	return f
}

// MustPoisson returns a random value from a Poisson distribution
// with parameter lambda using the generator's entropy source.
// It panics on error.
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
func (g *Generator) MustGamma(alpha, beta float64) float64 {
	f, err := g.Gamma(alpha, beta)
	if err != nil {
		panic(err)
	}
	return f
}
