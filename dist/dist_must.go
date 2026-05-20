//go:build randutil_must
// +build randutil_must

package dist

// MustBernoulli returns true with probability p and false otherwise.
// It panics on error.
func MustBernoulli(p float64) bool {
	return Default().MustBernoulli(p)
}

// MustCategorical samples an index in [0, len(weights)) with probability
// proportional to weights[i]. It panics on error.
func MustCategorical(weights []float64) int {
	return Default().MustCategorical(weights)
}

// MustExponential returns a random value from an exponential distribution
// with rate parameter lambda. It panics on error.
func MustExponential(lambda float64) float64 {
	return Default().MustExponential(lambda)
}

// MustNormal returns a random value from a normal distribution
// with mean mu and standard deviation sigma. It panics on error.
func MustNormal(mu, sigma float64) float64 {
	return Default().MustNormal(mu, sigma)
}

// MustUniform returns a random value from a uniform distribution
// in [min, max). It panics on error.
func MustUniform(minVal, maxVal float64) float64 {
	return Default().MustUniform(minVal, maxVal)
}

// MustPoisson returns a random value from a Poisson distribution
// with parameter lambda. It panics on error.
func MustPoisson(lambda float64) int {
	return Default().MustPoisson(lambda)
}

// MustGamma returns a random value from a gamma distribution
// with shape alpha and rate beta. It panics on error.
func MustGamma(alpha, beta float64) float64 {
	return Default().MustGamma(alpha, beta)
}
