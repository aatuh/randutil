//go:build randutil_must
// +build randutil_must

package uuid

// MustV4 returns a v4 UUID or panics on error.
func (g *Generator) MustV4() UUID {
	u, err := g.V4()
	if err != nil {
		panic(err)
	}
	return u
}

// MustV7 returns a v7 UUID or panics on error.
func (g *Generator) MustV7() UUID {
	u, err := g.V7()
	if err != nil {
		panic(err)
	}
	return u
}
