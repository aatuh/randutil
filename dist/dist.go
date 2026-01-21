package dist

import "time"

// Bernoulli returns true with probability p and false otherwise.
// p must be in [0,1].
func Bernoulli(p float64) (bool, error) {
	return Default().Bernoulli(p)
}

// Categorical samples an index in [0, len(weights)) with probability
// proportional to weights[i]. All weights must be >= 0 and at least
// one weight must be > 0.
func Categorical(weights []float64) (int, error) {
	return Default().Categorical(weights)
}

// Exponential samples an exponential(lambda). lambda must be > 0.
func Exponential(lambda float64) (float64, error) {
	return Default().Exponential(lambda)
}

// Normal returns a normal(mean, stddev) variate using Box-Muller.
// stddev must be >= 0. If stddev == 0, returns mean.
func Normal(mean, stddev float64) (float64, error) {
	return Default().Normal(mean, stddev)
}

// Uniform returns a random value from a uniform distribution in [min, max).
func Uniform(minVal, maxVal float64) (float64, error) {
	return Default().Uniform(minVal, maxVal)
}

// Poisson returns a random value from a Poisson distribution
// with parameter lambda.
func Poisson(lambda float64) (int, error) {
	return Default().Poisson(lambda)
}

// Gamma returns a random value from a gamma distribution
// with shape alpha and rate beta.
func Gamma(alpha, beta float64) (float64, error) {
	return Default().Gamma(alpha, beta)
}

// SeededClockNormal returns a normal variate around time.Now with
// jitter stddev seconds. It is a small example of composing dists.
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
