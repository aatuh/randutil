package uuid

type rng interface {
	Bytes(n int) ([]byte, error)
}
