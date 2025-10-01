package core

// Source is the entropy source abstraction used by generators. It matches
// io.Reader so callers can plug in crypto/rand.Reader, deterministic
// readers in tests, HSM-backed readers, etc.
type Source interface {
	Read(p []byte) (int, error)
}

// Default is the package-wide default entropy source used when a generator
// is constructed with a nil source. It dynamically proxies to the active
// core source (which defaults to crypto/rand.Reader).
var Default Source = Reader()
