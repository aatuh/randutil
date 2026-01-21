# randutil

CSPRNG-first random utilities for Go. Small, composable generators with
bias-free numeric helpers, strings/tokens, distributions, UUIDs, and
sampling utilities.

## Install

```bash
go get github.com/aatuh/randutil/v2
```

Requires Go 1.24+.

## Quick start

```go
package main

import (
  "fmt"

  "github.com/aatuh/randutil/v2"
)

func main() {
  r := randutil.Default()

  b := r.Numeric.MustBytes(16)
  tok := r.String.MustTokenURLSafe(24)
  u4 := r.UUID.MustV4()
  when := r.Time.MustDatetime()

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
  rng := core.New(adapters.DeterministicSource([]byte("seed")))
  gen := randstring.New(rng)

  s, _ := gen.String(12)
  fmt.Println(s)
}
```

For exact byte control in tests, pass a custom `io.Reader` into `core.New`.

## Security model

- Default entropy is `crypto/rand.Reader`.
- No process-wide mutable state; each generator binds its own source/RNG.
- Unbiased sampling (rejection sampling for ranges/charsets).
- Token string helpers return immutable strings; use `Token*Bytes` when secrets must be wipeable.
- Deterministic sources are for testing and benchmarks only.

If your source or RNG is not thread-safe, wrap it with
`adapters.LockedSource` or `adapters.LockedRNG`.

## Common recipes

URL-safe token:

```go
s := randstring.MustTokenURLSafe(24) // ~32 chars, URL-safe
```

Range and sampling:

```go
n := numeric.MustIntRange(10, 20) // inclusive
arr := []int{1, 2, 3, 4, 5}
collection.MustShuffle(arr)
subset := collection.MustSample(arr, 2)
```

UUIDs:

```go
u4 := uuid.MustV4()
u7 := uuid.MustV7()
```

Distributions:

```go
x := dist.MustNormal(0, 1)
k := dist.MustPoisson(12)
```

Email:

```go
mail := email.MustEmail(email.Options{TLD: "org"})
```

## Notes

Every `MustX(...)` has a non-panicking `X(...)` variant that returns
`(T, error)` for server/CLI contexts.

The `example_test.go` file contains executable examples that appear in
`godoc` and are run by `go test` to prevent documentation drift.
