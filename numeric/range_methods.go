package numeric

// IntRange returns a secure random int in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) IntRange(minInclusive int, maxInclusive int) (int, error) {
	return g.rng.IntRange(minInclusive, maxInclusive)
}

// AnyInt returns a secure random int from the full range of int.
func (g *Generator) AnyInt() (int, error) {
	return g.IntRange(minInt, maxInt)
}

// Int32Range returns a secure random int32 in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) Int32Range(minInclusive int32, maxInclusive int32) (int32, error) {
	return g.rng.Int32Range(minInclusive, maxInclusive)
}

// AnyInt32 returns a random int32 in the full int32 range.
func (g *Generator) AnyInt32() (int32, error) {
	return g.Int32Range(minInt32, maxInt32)
}

// PositiveInt32 returns a secure random positive int32.
func (g *Generator) PositiveInt32() (int32, error) {
	return g.Int32Range(1, maxInt32)
}

// NegativeInt32 returns a secure random negative int32.
func (g *Generator) NegativeInt32() (int32, error) {
	return g.Int32Range(minInt32, -1)
}

// Int64Range returns a secure random int64 in [minInclusive, maxInclusive]
// using the generator's entropy source.
func (g *Generator) Int64Range(minInclusive int64, maxInclusive int64) (int64, error) {
	return g.rng.Int64Range(minInclusive, maxInclusive)
}

// AnyInt64 returns a secure random int64 in the full int64 range.
func (g *Generator) AnyInt64() (int64, error) {
	return g.Int64Range(minInt64, maxInt64)
}

// PositiveInt64 returns a secure random positive int64.
func (g *Generator) PositiveInt64() (int64, error) {
	return g.Int64Range(1, maxInt64)
}

// NegativeInt64 returns a secure random negative int64.
func (g *Generator) NegativeInt64() (int64, error) {
	return g.Int64Range(minInt64, -1)
}
