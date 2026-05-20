//go:build randutil_must
// +build randutil_must

package numeric

// MustIntRange returns a secure random int in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustIntRange(minInclusive int, maxInclusive int) int {
	v, err := g.IntRange(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
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

// MustInt32Range returns a secure random int32 in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustInt32Range(minInclusive int32, maxInclusive int32) int32 {
	v, err := g.Int32Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
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

// MustPositiveInt32 returns a secure random positive int32.
// It panics on error.
func (g *Generator) MustPositiveInt32() int32 {
	v, err := g.PositiveInt32()
	if err != nil {
		panic(err)
	}
	return v
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

// MustInt64Range returns a secure random int64 in [minInclusive, maxInclusive].
// It panics on error.
func (g *Generator) MustInt64Range(minInclusive int64, maxInclusive int64) int64 {
	v, err := g.Int64Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return v
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

// MustPositiveInt64 returns a secure random positive int64.
// It panics on error.
func (g *Generator) MustPositiveInt64() int64 {
	v, err := g.PositiveInt64()
	if err != nil {
		panic(err)
	}
	return v
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
