// Package randutil provides cryptographically secure random utilities
// organized into focused subpackages. Generators accept an injectable
// RNG and default to crypto/rand.Reader:
//
//   - core: Basic random number generation primitives and entropy source
//     ports
//   - adapters: Entropy source adapters (crypto, deterministic, locking)
//   - numeric: Random number generation for various numeric types and booleans
//   - randstring: Random string generation and token creation
//   - dist: Statistical distributions
//   - email: Random email address generation with customizable options
//   - collection: Random sampling, shuffling, and weighted selection for slices
//   - randtime: Random datetime generation functions
//   - uuid: UUID generation
package randutil
