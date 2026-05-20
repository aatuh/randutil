//go:build randutil_must
// +build randutil_must

package numeric

// MustUint64 returns a random uint64 from crypto/rand. It panics on error.
//
// Returns:
//   - uint64: A random uint64 value.
func MustUint64() uint64 {
	u, err := Default().Uint64()
	if err != nil {
		panic(err)
	}
	return u
}

// MustUint64 returns a random uint64 from the generator. It panics on error.
func (g *Generator) MustUint64() uint64 {
	u, err := g.Uint64()
	if err != nil {
		panic(err)
	}
	return u
}

// MustUint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - uint64: A random uint64 in [0, n).
func MustUint64n(n uint64) uint64 {
	u, err := Default().Uint64n(n)
	if err != nil {
		panic(err)
	}
	return u
}

// MustUint64n returns a random uint64 in [0, n) from the generator.
func (g *Generator) MustUint64n(n uint64) uint64 {
	u, err := g.Uint64n(n)
	if err != nil {
		panic(err)
	}
	return u
}

// MustIntn returns a uniform random int in [0, n). It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - int: A random int in [0, n).
func MustIntn(n int) int {
	v, err := Default().Intn(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustIntn returns a random int in [0, n) from the generator.
func (g *Generator) MustIntn(n int) int {
	v, err := g.Intn(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustInt64n returns a uniform random int64 in [0, n). It panics on error.
//
// Parameters:
//   - n: The upper bound (exclusive) for the random number.
//
// Returns:
//   - int64: A random int64 in [0, n).
func MustInt64n(n int64) int64 {
	v, err := Default().Int64n(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustInt64n returns a random int64 in [0, n) from the generator.
func (g *Generator) MustInt64n(n int64) int64 {
	v, err := g.Int64n(n)
	if err != nil {
		panic(err)
	}
	return v
}

// MustFloat64 returns a uniform random float64 in [0.0, 1.0) with 53 bits
// of precision. It panics on error.
//
// Returns:
//   - float64: A random float64 in [0.0, 1.0).
func MustFloat64() float64 {
	v, err := Default().Float64()
	if err != nil {
		panic(err)
	}
	return v
}

// MustFloat64 returns a uniform random float64 in [0,1) from the generator.
func (g *Generator) MustFloat64() float64 {
	v, err := g.Float64()
	if err != nil {
		panic(err)
	}
	return v
}
