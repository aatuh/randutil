// Package randutil provides cryptographically secure random utilities
// organized into focused subpackages. Generators accept an injectable
// RNG and default to crypto/rand.Reader. Generators are concurrency-safe
// iff the injected RNG is; crypto/rand.Reader is safe for concurrent use.
//
//   - core: Basic random number generation primitives and entropy source
//     ports
//   - adapters: Entropy source adapters (crypto, deterministic, derived, locking)
//   - numeric: Random number generation for various numeric types and booleans
//   - randstring: Random string generation and token creation
//   - dist: Statistical distributions
//   - email: Random email address generation with customizable options
//   - collection: Random sampling, shuffling, and weighted selection for slices
//   - randtime: Random datetime generation functions
//   - nanoid: NanoID-style identifiers
//   - ulid: ULID identifiers
//   - uuid: UUID generation
//
// The Workspace API provides domain-separated streams derived from a root
// seed for stable tests and reduced cross-stream coupling. Use Fast when you
// want a derived high-throughput CSPRNG, and Default/Secure for OS RNG.
package randutil
