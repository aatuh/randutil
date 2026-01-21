package core

// Source is the entropy source abstraction used by generators. It matches
// io.Reader so callers can plug in crypto/rand.Reader, deterministic
// readers in tests, HSM-backed readers, etc.
type Source interface {
	Read(p []byte) (int, error)
}

// RNG is the randomness port used across the repo. Implementations should be
// safe for concurrent use if the underlying Source is safe.
type RNG interface {
	Read(p []byte) (int, error)
	Fill(p []byte) error
	Bytes(n int) ([]byte, error)
	Uint64() (uint64, error)
	Uint64n(n uint64) (uint64, error)
	Intn(n int) (int, error)
	Int64n(n int64) (int64, error)
	IntRange(minInclusive, maxInclusive int) (int, error)
	Int32Range(minInclusive, maxInclusive int32) (int32, error)
	Int64Range(minInclusive, maxInclusive int64) (int64, error)
	Float64() (float64, error)
	Bool() (bool, error)
}
