package randtime

type rng interface {
	IntRange(minInclusive, maxInclusive int) (int, error)
	Float64() (float64, error)
}
