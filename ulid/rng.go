package ulid

type rng interface {
	Fill(p []byte) error
}
