// Package randutil provides cryptographically secure random number generation
// and related utilities organized into focused sub-packages:
//
//   - core: Basic random number generation primitives and entropy source
//     management
//   - numeric: Random number generation for various numeric types and booleans
//   - randstring: Random string generation and token creation
//   - email: Random email address generation with customizable options
//   - collection: Random sampling, shuffling, and weighted selection for slices
//   - randtime: Random datetime generation functions
//   - uuid: UUID generation
//
// All functions use crypto/rand by default for cryptographically secure
// randomness.
package randutil
