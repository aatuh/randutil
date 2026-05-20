package numeric

type rng interface {
	Bytes(n int) ([]byte, error)
	Fill(p []byte) error
	Uint64() (uint64, error)
	Uint64n(n uint64) (uint64, error)
	Intn(n int) (int, error)
	Int64n(n int64) (int64, error)
	Float64() (float64, error)
	Bool() (bool, error)
	IntRange(minInclusive, maxInclusive int) (int, error)
	Int32Range(minInclusive, maxInclusive int32) (int32, error)
	Int64Range(minInclusive, maxInclusive int64) (int64, error)
}
