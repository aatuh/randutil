# Release Process

`randutil` is a public Go module at `github.com/aatuh/randutil/v2`.
Maintain SemVer compatibility for the `v2` API: patch releases fix bugs and
docs, minor releases add backward-compatible API, and breaking changes require a
new major module path.

## Compatibility

- Public Go floor: `go 1.25.0` from `go.mod`.
- CI uses `actions/setup-go` with `go-version-file: go.mod`.
- Local audit toolchain observed during this update: `go1.26.3-X:nodwarf5`.
- The Go floor was raised from 1.24.0 because `golang.org/x/crypto v0.51.0`
  declares `go 1.25.0`.
- Newer Go versions are expected to work unless CI or local checks show a
  regression; do not claim support for untested older Go versions.

## Required Checks

Run the local release gate before tagging:

```bash
make finalize
```

That gate installs pinned tools, formats, vets, lints, checks vulnerabilities,
runs gosec, tidies both modules, runs normal/build-tag/race tests, runs fuzz
smoke, and cleans the test cache.

GitHub release readiness requires the tracked workflows to pass:

- Go CI test job
- dependency review on pull requests
- CodeQL
- Scorecard
- govulncheck SARIF upload and failure gate
- gosec SARIF upload and failure gate
- race tests with coverage
- `randutil_ci` and `randutil_must` build-tag tests
- fuzz smoke

Branch protection and required checks are configured in GitHub repository
settings; they are not enforceable from this repository alone.

## Checklist

1. Confirm `CHANGELOG.md` has an entry for the release.
2. Run `make finalize`.
3. Confirm GitHub Actions are green for the release commit.
4. Check GitHub Security Advisories for open private reports.
5. Tag releases as `vMAJOR.MINOR.PATCH`, for example `v2.1.2`.
6. Publish release notes with the changelog entry, CI result link, and benchmark
   summary when performance-sensitive code changed.
