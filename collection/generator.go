package collection

import "github.com/aatuh/randutil/v2/core"

// Generator builds collection-related random operations using a core RNG.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator[T any] struct {
	rng core.RNG
}

// New returns a collection Generator. If rng is nil, crypto/rand is used.
func New[T any](rng core.RNG) *Generator[T] {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator[T]{rng: rng}
}

// NewWithSource returns a collection Generator bound to src.
func NewWithSource[T any](src core.Source) *Generator[T] {
	return New[T](core.New(src))
}

var defaultRNG = core.New(nil)

// Default returns the package-wide default generator.
func Default[T any]() *Generator[T] {
	return &Generator[T]{rng: defaultRNG}
}

// Shuffle performs an in-place secure Fisher-Yates shuffle of the slice.
func (g *Generator[T]) Shuffle(slice []T) error {
	return shuffleWithRNG(g.rngOrDefault(), slice)
}

// Sample returns k items uniformly at random from s without replacement.
func (g *Generator[T]) Sample(s []T, k int) ([]T, error) {
	return sampleWithRNG(g.rngOrDefault(), s, k)
}

// PickOne returns one random element from the slice.
func (g *Generator[T]) PickOne(slice []T) (T, error) {
	return pickOneWithRNG(g.rngOrDefault(), slice)
}

// Perm returns a shuffled copy of slice.
func (g *Generator[T]) Perm(slice []T) ([]T, error) {
	dup := make([]T, len(slice))
	copy(dup, slice)
	if err := g.Shuffle(dup); err != nil {
		return nil, err
	}
	return dup, nil
}

// WeightedChoice returns one item where probability is proportional to
// its non-negative weight. Zero-weight items are never chosen.
func (g *Generator[T]) WeightedChoice(items []T, weights []float64) (T, error) {
	return weightedChoiceWithRNG(g.rngOrDefault(), items, weights)
}

// WeightedSample returns k items where probability is proportional to
// their non-negative weights. Zero-weight items are never chosen.
func (g *Generator[T]) WeightedSample(items []T, weights []float64, k int) ([]T, error) {
	return weightedSampleWithRNG(g.rngOrDefault(), items, weights, k)
}

// PickByProbability returns items independently with probability p in [0,1].
func (g *Generator[T]) PickByProbability(xs []T, p float64) ([]T, error) {
	return pickByProbabilityWithRNG(g.rngOrDefault(), xs, p)
}

func (g *Generator[T]) rngOrDefault() core.RNG {
	if g == nil || g.rng == nil {
		return defaultRNG
	}
	return g.rng
}
