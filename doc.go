// Package randutil provides cryptographically secure random number generation
// and related utilities organized into focused sub-packages:
//
//   - core: Basic random number generation primitives and entropy source management
//   - numeric: Random number generation for various numeric types and booleans
//   - string: Random string generation, token creation, and email addresses
//   - collection: Random sampling, shuffling, and weighted selection for slices
//   - time: Random datetime generation functions
//   - dist: Probability distributions (existing sub-package)
//   - uuid: UUID generation (existing sub-package)
//
// All functions use crypto/rand by default for cryptographically secure randomness.
package randutil
