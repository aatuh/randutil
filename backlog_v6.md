Here’s what I’d change in **randutil** right now to push it toward “boringly excellent”: secure-by-default, reproducible, easy to test, and CI that doesn’t drift.

---

## What I noticed in your current repo

* **Go version mismatch**: `go.mod` is `go 1.24.0`, but README still claims `Go 1.20+`, and Dependabot’s `go-version` is `1.20.x`. That’s confusing for users and for automation.
* **Workflows use `@main` and `@latest`** in multiple places:

  * `uses: actions/checkout@main`, `actions/setup-go@main`, `github/codeql-action@main`, `ossf/scorecard-action@main`
  * tools installed via `go install ...@latest`
    This is the #1 reason CI becomes non-reproducible (and occasionally breaks on a random Tuesday).
* You’re already doing a lot right: CodeQL + Scorecard + gosec + govulncheck + golangci-lint is a solid baseline.

---

## Version pinning: best versions to use (as of **2026-01-23**)

### GitHub Actions (pin these)

* `actions/checkout@v5.0.0` (Node 24 runtime; requires runner >= `v2.327.1`) ([GitHub][1])
* `actions/setup-go@v6.2.0` (current stable; supports `toolchain` directive handling in `go.mod`) ([GitHub][2])
* `github/codeql-action/*@v4.31.10` (use `init`, `autobuild`, `analyze`, and `upload-sarif` consistently) ([GitHub][3])
* `ossf/scorecard-action@v2.4.3` ([GitHub][4])

**Runner pinning**

* Replace `runs-on: ubuntu-latest` with a concrete image like `ubuntu-24.04` (because `ubuntu-latest` can and does change underneath you).

**Add two security-hardening actions (high ROI)**

* `step-security/harden-runner@v2.14.0`
* `actions/dependency-review-action@v4.8.2`

### Security/QA tooling (stop installing `@latest`)

Pin these in CI:

* `golangci-lint` → `github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.8.0` ([GitHub][5])
* `gosec` → `github.com/securego/gosec/v2/cmd/gosec@v2.22.10` ([GitHub][6])
* `govulncheck` → `golang.org/x/vuln/cmd/govulncheck@v1.1.4` ([GitHub][7])

### The cleanest way to pin tool versions in Go 1.24

Go 1.24 adds first-class **tool dependencies** in `go.mod`, so you can pin linters/scanners like real deps (instead of “some CI step installs whatever”). ([Go][8])

That’s the modern, developer-friendly approach for a library like this.

---

## “Randomness workspace”: what problem it solves (and how)

The core problem it solves is **randomness coupling**:

* In real systems, “randomness” gets consumed by many features: IDs, tokens, shuffles, sampling, backoff jitter, etc.
* If everything pulls from one global source, the sequence becomes **order-dependent**.

  * Tests become flaky: small refactors change call ordering => different randomness.
  * Security reviews get harder: you can’t easily guarantee “this nonce stream can’t be influenced by that feature”.
  * Deterministic simulations become painful.

A workspace solves it by giving you **domain-separated sub-streams**, derived from a single root secret:

* `Stream("ids")`, `Stream("tokens")`, `Stream("nonces")` are independent.
* Each stream is derived via HKDF (domain separation by label) and then expanded by a stream cipher PRNG.
* That yields:

  * **Reproducibility** when seeded deterministically (great for tests/sims)
  * **Isolation** in production (one consumer can’t “pull forward” another consumer’s entropy stream)

It’s basically: “structured randomness, not a free-for-all global RNG”.

---

## Comprehensive backlog

### EPIC A — CI reproducibility + supply chain hardening (P0)

1. **Pin all GitHub Actions to stable versions**

   * Replace all `@main` with pinned tags:

     * checkout `v5.0.0`, setup-go `v6.2.0`, codeql `v4.31.10`, scorecard `v2.4.3`. ([GitHub][1])
2. **Pin runner image**

   * `ubuntu-latest` → `ubuntu-24.04`.
3. **Add `harden-runner` at the top of each workflow**

   * Use egress-auditing mode first; then tighten allowed endpoints if you want.
4. **Add Dependency Review on PRs**

   * Block newly introduced vulnerable deps and risky licenses.
5. **Stop `go install ...@latest`**

   * Pin `golangci-lint`, `gosec`, `govulncheck` versions (see above). ([GitHub][5])
6. **Make all scanners emit SARIF and upload**

   * `gosec -fmt=sarif` → upload-sarif
   * `govulncheck -format sarif` (supported in v1.1.x line) → upload-sarif ([GitHub][7])
7. **Set default workflow permissions to read-only**

   * Put `permissions: {}` at workflow root, grant only per-job what’s needed.

### EPIC B — Go 1.24 alignment + toolchain discipline (P0)

8. **Update README to state Go 1.24+ (or exactly 1.24.0)**
9. **Update Dependabot `go-version` to 1.24.0**
10. **Add `toolchain go1.24.0` in `go.mod`**

* Makes builds more reproducible; setup-go v6 respects it. ([GitHub][9])

11. **Adopt Go 1.24 tool dependencies**

* Pin tool versions in `go.mod` (preferred over ad-hoc CI installs). ([Go][8])

### EPIC C — Workspace security + lifecycle (P0/P1)

12. **Add `Workspace.Close()`**

* Zero the root seed and evict/zero derived stream state where possible.

13. **Add hierarchical workspaces**

* `w.Sub("payments").Stream("nonces")` style: reduces label collision risks and improves ergonomics.

14. **Eliminate panic paths in derivation**

* Anything that can currently `panic` should either:

  * return an error, or
  * be a clearly named `Must...` API.
    Panics are denial-of-service when a library is used in services.

15. **Document hard “DO NOTs”**

* e.g. deterministic sources must not be used for secrets unless the seed is secret and protected.

### EPIC D — Secure “math/rand conveniences” (P1)

These are surprisingly demanded, because people want safe versions of common helpers:
16. `Shuffle[T any]`, `Perm(n)`, `Choose[T]`, `Sample[T](k)`
17. `WeightedChoice` (with stable/explicit error handling)
18. `Jitter(duration, pct)` for backoff (CSPRNG-based)
19. Provide both:

* “fast-ish CSPRNG” (your derived ChaCha stream)
* “system CSPRNG” (`crypto/rand.Reader`)
  so callers can choose by threat model.

### EPIC E — Test strategy that proves properties (P1)

20. **Property tests** for uniform sampling / modulo-bias avoidance
21. **Concurrency tests** for Workspace stream caching (race-enabled)
22. **Golden vectors** for deterministic modes (so refactors can’t silently change output)
23. **Fuzz targets** for parse/format functions (UUID/ULID/NanoID)
24. Benchmarks:

* `crypto/rand` vs derived stream
* ID generation throughput

### EPIC F — Public API polish (P1/P2)

25. Introduce small interfaces for injection:

* e.g. `type Workspace interface { Stream(label string) (core.Source, error) }`

26. Provide a “safe defaults” constructor and a “fully explicit” constructor
27. Ensure every exported symbol has:

* threat-model note (when relevant)
* deterministic vs non-deterministic behavior spelled out

### EPIC G — “Missing feature” bets that could make randutil stand out (P2)

These are uncommon in existing libs and can become your differentiators:
28. **Randomness recorder/replayer**

* Wrap a `Source` to record bytes, replay them later: perfect for debugging, simulations, deterministic CI failures.

29. **Stream usage accounting**

* Count bytes drawn per label; optionally expose metrics hooks.

30. **“Policy mode”**

* A build tag / option that forbids insecure patterns (e.g. disallow deterministic seeds in non-test builds).

---

## If you want one concrete “next commit”

Do this first (highest leverage):

1. Pin Actions (drop `@main`), pin runner image, pin Go tooling versions
2. Align Go versions everywhere (README + dependabot + CI), and move tool pins into Go 1.24 tool deps

That alone will make the repo feel dramatically more “serious” and predictable.

If you paste your current `.github/workflows/*.yml` goals (e.g., do you want SARIF for gosec/govulncheck, do you want PR gating), I can give you an exact cleaned-up workflow set that matches the backlog above.

[1]: https://github.com/actions/checkout/releases "https://github.com/actions/checkout/releases"
[2]: https://github.com/actions/setup-go/releases "https://github.com/actions/setup-go/releases"
[3]: https://github.com/github/codeql-action/releases "https://github.com/github/codeql-action/releases"
[4]: https://github.com/ossf/scorecard-action/releases "https://github.com/ossf/scorecard-action/releases"
[5]: https://github.com/golangci/golangci-lint/releases "https://github.com/golangci/golangci-lint/releases"
[6]: https://github.com/securego/gosec/releases "https://github.com/securego/gosec/releases"
[7]: https://github.com/golang/vuln/releases "https://github.com/golang/vuln/releases"
[8]: https://go.dev/doc/go1.24 "https://go.dev/doc/go1.24"
[9]: https://github.com/actions/setup-go "https://github.com/actions/setup-go"
