# Benchmarks

Use benchmarks to detect performance regressions in random generation,
encoding, sampling, distributions, and identifier helpers. Do not make
performance claims without recording the environment and command output.

Run all benchmarks with:

```bash
make bench
```

The target runs:

```bash
go test -run=^$ -bench=. -benchmem ./...
```

For release notes, record:

- Go version from `go version`
- OS and architecture
- CPU model if available
- command used
- package-level benchmark output relevant to the change
- comparison baseline, usually the previous release tag

If a release changes buffering, derived streams, numeric sampling, string/token
encoding, or identifier generation, include a short benchmark summary in the
release notes even when results are neutral.
