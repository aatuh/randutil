package collection

type rng interface {
	Uint64n(n uint64) (uint64, error)
	Intn(n int) (int, error)
	Float64() (float64, error)
}
