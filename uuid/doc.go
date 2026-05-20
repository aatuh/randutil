// Package uuid provides RFC 4122 v4 (random) and RFC 9562 v7
// (time-ordered) UUID helpers built on randutil.
//
// UUID v7 values encode Unix milliseconds for ordering, but this package does
// not make UUID v7 values monotonic within the same millisecond. Generators are
// concurrency-safe iff the injected RNG is safe.
package uuid
