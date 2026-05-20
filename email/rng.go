package email

type rng interface {
	Bytes(n int) ([]byte, error)
	Uint64n(n uint64) (uint64, error)
	Fill(p []byte) error
	IntRange(minInclusive, maxInclusive int) (int, error)
}
