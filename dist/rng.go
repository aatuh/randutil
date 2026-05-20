package dist

type rng interface {
	Float64() (float64, error)
}
