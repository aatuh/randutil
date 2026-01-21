package numeric

// IntRange returns a secure random int in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) IntRange(minInclusive int, maxInclusive int) (int, error) {
	return g.rng.IntRange(minInclusive, maxInclusive)
}

// MustIntRange returns a secure random int in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustIntRange(minInclusive int, maxInclusive int) int {
	v, err := g.IntRange(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
}

// AnyInt returns a secure random int from the full range of int.
func (g *Generator) AnyInt() (int, error) {
	return g.IntRange(minInt, maxInt)
}

// MustAnyInt returns a secure random int from the full range of int.
// It panics on error.
func (g *Generator) MustAnyInt() int {
	v, err := g.AnyInt()
	if err != nil {
		panic(err)
	}
	return v
}

// Int32Range returns a secure random int32 in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) Int32Range(minInclusive int32, maxInclusive int32) (int32, error) {
	return g.rng.Int32Range(minInclusive, maxInclusive)
}

// MustInt32Range returns a secure random int32 in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustInt32Range(minInclusive int32, maxInclusive int32) int32 {
	v, err := g.Int32Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
}

// AnyInt32 returns a random int32 in the full int32 range.
func (g *Generator) AnyInt32() (int32, error) {
	return g.Int32Range(minInt32, maxInt32)
}

// MustAnyInt32 returns a random int32 in the full int32 range.
// It panics on error.
func (g *Generator) MustAnyInt32() int32 {
	v, err := g.AnyInt32()
	if err != nil {
		panic(err)
	}
	return v
}

// PositiveInt32 returns a secure random positive int32.
func (g *Generator) PositiveInt32() (int32, error) {
	return g.Int32Range(1, maxInt32)
}

// MustPositiveInt32 returns a secure random positive int32.
// It panics on error.
func (g *Generator) MustPositiveInt32() int32 {
	v, err := g.PositiveInt32()
	if err != nil {
		panic(err)
	}
	return v
}

// NegativeInt32 returns a secure random negative int32.
func (g *Generator) NegativeInt32() (int32, error) {
	return g.Int32Range(minInt32, -1)
}

// MustNegativeInt32 returns a secure random negative int32.
// It panics on error.
func (g *Generator) MustNegativeInt32() int32 {
	v, err := g.NegativeInt32()
	if err != nil {
		panic(err)
	}
	return v
}

// Int64Range returns a secure random int64 in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) Int64Range(minInclusive int64, maxInclusive int64) (int64, error) {
	return g.rng.Int64Range(minInclusive, maxInclusive)
}

// MustInt64Range returns a secure random int64 in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustInt64Range(minInclusive int64, maxInclusive int64) int64 {
	v, err := g.Int64Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
}

// AnyInt64 returns a secure random int64 in the full int64 range.
func (g *Generator) AnyInt64() (int64, error) {
	return g.Int64Range(minInt64, maxInt64)
}

// MustAnyInt64 returns a secure random int64 in the full int64 range.
// It panics on error.
func (g *Generator) MustAnyInt64() int64 {
	v, err := g.AnyInt64()
	if err != nil {
		panic(err)
	}
	return v
}

// PositiveInt64 returns a secure random positive int64.
func (g *Generator) PositiveInt64() (int64, error) {
	return g.Int64Range(1, maxInt64)
}

// MustPositiveInt64 returns a secure random positive int64.
// It panics on error.
func (g *Generator) MustPositiveInt64() int64 {
	v, err := g.PositiveInt64()
	if err != nil {
		panic(err)
	}
	return v
}

// NegativeInt64 returns a secure random negative int64.
func (g *Generator) NegativeInt64() (int64, error) {
	return g.Int64Range(minInt64, -1)
}

// MustNegativeInt64 returns a secure random negative int64.
// It panics on error.
func (g *Generator) MustNegativeInt64() int64 {
	v, err := g.NegativeInt64()
	if err != nil {
		panic(err)
	}
	return v
}
