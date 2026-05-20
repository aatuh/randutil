// Package ulid provides ULID generation helpers (Universally Unique
// Lexicographically Sortable Identifiers).
//
// ULIDs encode time for lexical ordering, but this package does not make ULIDs
// monotonic within the same millisecond.
// Generators are concurrency-safe iff the injected RNG is safe.
package ulid
