# randutil

CSPRNG-first random utilities for Go. Small, composable generators with
bias-free numeric helpers, strings/tokens, distributions, UUIDs, ULIDs,
NanoIDs, and sampling utilities.

## Install

```bash
go get github.com/aatuh/randutil/v2
```

Requires Go 1.24.0+.

## Quick start

```go
package main

import (
  "fmt"
  "log"

  "github.com/aatuh/randutil/v2"
)

func main() {
  r := randutil.Default()

  b, err := r.Numeric.Bytes(16)
  if err != nil {
    log.Fatal(err)
  }
  tok, err := r.String.TokenURLSafe(24)
  if err != nil {
    log.Fatal(err)
  }
  u4, err := r.UUID.V4()
  if err != nil {
    log.Fatal(err)
  }
  when, err := r.Time.Datetime()
  if err != nil {
    log.Fatal(err)
  }

  fmt.Println(len(b), tok, u4, when)
}
```

## Deterministic testing

Use a deterministic source and pass it into `core.New`, then share the RNG
across generators:

```go
package main

import (
  "fmt"

  "github.com/aatuh/randutil/v2/adapters"
  "github.com/aatuh/randutil/v2/core"
  "github.com/aatuh/randutil/v2/randstring"
)

func main() {
  src, err := adapters.DeterministicSource([]byte("seed"))
  if err != nil {
    panic(err)
  }
  rng := core.New(src)
  gen := randstring.New(rng)

  s, _ := gen.String(12)
  fmt.Println(s)
}
```

For exact byte control in tests, pass a custom `io.Reader` into `core.New`.
If you want the intent to be explicit, use `adapters/deterministic`.
Deterministic sources are for tests and benchmarks only; DO NOT USE FOR
TOKENS / AUTH unless the seed is high-entropy and kept secret.

## Workspace and domain separation

```go
ws := randutil.NewWorkspace(randutil.SecureRoot())
sessions, _ := ws.Rand("sessions")
tok, _ := sessions.String.TokenURLSafe(24)
fmt.Println(tok)
```

For deterministic fixtures:

```go
ws := randutil.NewWorkspace(randutil.DeterministicRoot([]byte("seed")))
sampling, _ := ws.Rand("sampling")
v, _ := sampling.Numeric.IntRange(1, 10)
fmt.Println(v)
```

Sub-workspaces reduce label collisions:

```go
billing, _ := ws.Sub("billing")
nonces, _ := billing.Rand("nonces")
nonce, _ := nonces.String.TokenURLSafe(24)
fmt.Println(nonce)
```

Workspaces can also track bytes read per cached stream:

```go
ws := randutil.NewWorkspaceWithOptions(randutil.SecureRoot(), randutil.WorkspaceOptions{
  UsageHook: func(label string, delta uint64) {
    fmt.Println(label, delta)
  },
})
tokens, _ := ws.Rand("tokens")
_, _ = tokens.String.TokenURLSafe(24)
used, _ := ws.Usage("tokens")
fmt.Println(used)
```

For a single derived stream without a workspace:

```go
r, err := randutil.Derive([]byte("seed"), "payments")
if err != nil {
  panic(err)
}
id, _ := r.UUID.V7()
fmt.Println(id)
```

Workspace streams are derived via HKDF-SHA256 + ChaCha20; for strict FIPS/OS
RNG compliance, use `crypto/rand.Reader` directly.

For a fast derived CSPRNG seeded from `crypto/rand`:

```go
fast, _ := randutil.Fast()
id, _ := fast.UUID.V7()
fmt.Println(id)
```

## Must helpers (opt-in)

`Must*` helpers are gated behind the build tag `randutil_must` to avoid
accidental panics in production. Enable the tag if you want them:

```bash
go test -tags=randutil_must ./...
go build -tags=randutil_must ./...
```

## Security model

- Default entropy is `crypto/rand.Reader`.
- No process-wide mutable state; each generator binds its own source/RNG.
- Generators are concurrency-safe iff the injected RNG is; `crypto/rand.Reader`
  is safe for concurrent use.
- For high-throughput workloads, wrap sources with `adapters.BufferedSource`
  to amortize small reads.
- Unbiased sampling (rejection sampling for ranges/charsets).
- Token string helpers return immutable strings; use `Token*Bytes` when secrets must be wipeable.
- Deterministic sources are for testing and benchmarks only.
- Build with `-tags=randutil_policy` to make deterministic source/root constructors fail with `ErrDeterministicDisabled`.

If your source or RNG is not thread-safe, wrap it with
`adapters.LockedSource` or `adapters.LockedRNG`.

## Common recipes

URL-safe token:

```go
s, _ := randstring.TokenURLSafe(24) // ~32 chars, URL-safe
```

Range and sampling:

```go
n, _ := numeric.IntRange(10, 20) // inclusive
arr := []int{1, 2, 3, 4, 5}
_ = collection.Shuffle(arr)
subset, _ := collection.Sample(arr, 2)
```

UUIDs:

```go
u4, _ := uuid.V4()
u7, _ := uuid.V7()
```

ULID / NanoID:

```go
u, _ := ulid.ID()
id, _ := nanoid.ID()
```

Distributions:

```go
x, _ := dist.Normal(0, 1)
k, _ := dist.Poisson(12)
```

Email:

```go
mail, _ := email.Email(email.Options{TLD: "org"})
```

Record and replay entropy for debugging deterministic failures:

```go
src := adapters.NewRecorder(adapters.CryptoSource())
rng := core.New(src)
_, _ = rng.Bytes(16)
replay := src.Replay()
_ = replay
```

## Notes

Every `MustX(...)` has a non-panicking `X(...)` variant that returns
`(T, error)` for server/CLI contexts.

The `example_test.go` file contains executable examples that appear in
`godoc` and are run by `go test` to prevent documentation drift.
