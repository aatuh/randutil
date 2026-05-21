# Changelog

All notable changes to `randutil` are recorded here.

This project follows Semantic Versioning for the public `v2` module API.

## Unreleased

No unreleased changes yet.

## v2.1.2 - 2026-05-21

### Added

- Production-readiness documentation for releases, Go compatibility, required
  checks, and benchmark reporting.
- `make bench` for repeatable local benchmark collection.
- Root facade tests covering default constructors, derived RNGs, fast streams,
  and collection composition.

### Fixed

- Reject `NaN` and infinite probabilities in `collection.PickByProbability`
  with `core.ErrInvalidProbability`.
- Allow full signed integer ranges to return `math.MinInt64` and 64-bit
  `math.MinInt` without `ErrResultOutOfRange`.
- Reject UUID v7 timestamps that exceed the 48-bit RFC 9562 timestamp field.

### Changed

- Updated runtime dependency `golang.org/x/crypto` to `v0.51.0`, which raises
  the module Go floor to `1.25.0`.
- Refreshed pinned local audit tools and GitHub Actions used by CI, CodeQL, and
  Scorecard.

### Documentation

- Added production-boundary guidance for FIPS/audited-crypto expectations, seed
  quality, and direct `crypto/rand.Reader` use.
- Clarified that UUID v7 and ULID values are time-ordered but not monotonic
  within the same millisecond.
- Recorded the `v2.1.2` benchmark baseline in
  `docs/benchmarks-v2.1.2.md`.

## v2.1.1 and Earlier

See the Git tag history for released changes before this changelog was added.
