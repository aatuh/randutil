# randutil

CSPRNG-first random utilities for Go. One small package that does the
common random things: bias-free integers, bytes, strings/tokens, emails,
shuffle/sample for slices, and realistic date/time helpers.

## Install

```bash
go get github.com/aatuh/randutil/v2
```

## Quick start

```go
package main

import (
  "fmt"

  "github.com/aatuh/randutil/v2/numeric"
  "github.com/aatuh/randutil/v2/collection"
  "github.com/aatuh/randutil/v2/randstring"
  "github.com/aatuh/randutil/v2/email"
  "github.com/aatuh/randutil/v2/randtime"
  "github.com/aatuh/randutil/v2/uuid"
)

func main() {
  // Numbers
  n := numeric.MustIntRange(10, 20)     // inclusive
  f := numeric.MustFloat64()            // [0,1)
  b := numeric.MustBytes(16)            // 16 random bytes
  ok := numeric.MustBool()

  // Strings / tokens
  s   := randstring.MustString(12)                  // lower-case alnum
  hex := randstring.MustHex(32)                     // 32 hex chars (16 bytes)
  b64 := randstring.MustBase64(24)                  // encodes 24 random bytes
  tok := randstring.MustTokenURLSafe(24)            // URL-safe base64
  
  // Email addresses
  mail := email.MustEmailSimple(16)                 // exact-length local@domain.com
  mail2 := email.MustEmail(email.EmailOptions{TLD: "org"}) // options

  // Collections
  arr := []int{1,2,3,4,5}
  collection.MustShuffle(arr)                       // in-place Fisher–Yates
  top2 := collection.MustSample(arr, 2)             // k without replacement
  pick := collection.MustSlicePickOne(arr)
  pickedMany := collection.MustPickByProbability([]string{"a","b","c"}, 0.5)

  // Time
  t  := randtime.MustDatetime()
  p  := randtime.MustTimeInNearPast()
  fu := randtime.MustTimeInNearFuture()

  // UUIDs
  u4 := uuid.MustV4()
  u7 := uuid.MustV7()

  fmt.Println(n, f, len(b), ok, s, hex, b64, tok, mail, mail2, arr, top2, pick, pickedMany, t, p, fu, u4, u7)
}
```

Note: Every `MustX(...)` has a non-panicking `X(...)` variant that returns
`(T, error)` for use in servers/CLIs where you want to handle errors.

## API overview

### Numbers

Uniform integers without modulo bias (rejection sampling under the hood),
`[0,1)` float64, and byte helpers. Range functions are inclusive.

* `Uint64()`
* `Uint64n(n uint64)` → \[0,n)
* `Intn(n int)`, `Int64n(n int64)` → \[0,n)
* `IntRange(min, max)`, `Int32Range`, `Int64Range` (inclusive)
* `AnyInt*`, `Positive*`, `Negative*`
* `Float64()`
* `Bytes(n)`, `Fill(b)`

### Booleans

* `Bool()` / `MustBool()`

### Strings & tokens

Lower-case alnum by default; hex/base64 helpers.

* `String(n)` and `StringWithCharset(n, charset)`
* `Hex(len)` where `len` must be even (2 chars per byte)
* `Base64(byteLen)`, `TokenHex(byteLen)`, `TokenBase64(byteLen)`,
  `TokenURLSafe(byteLen)`

### Email addresses

Random email generation with customizable local parts, domains, and TLDs.

* `Email(opts)` - Full control with EmailOptions
* `EmailSimple(totalLength)` - Legacy exact-length emails
* `EmailWithCustomLocal(localPart)`
* `EmailWithCustomDomain(domainPart)`
* `EmailWithCustomTLD(tld)`
* `EmailWithRandomTLD()`
* `EmailWithoutTLD()`

### Collections (slices)

Unbiased Fisher–Yates shuffle, sampling without replacement, and simple picks.

* `Shuffle[T]([]T)`
* `Sample[T]([]T, k)` – returns a new slice, no duplicates
* `Perm(n)` – random permutation of 0..n-1
* `SlicePickOne[T]([]T)`
* `PickByProbability[T]([]T, p float64)` – picks each item with prob `p`
  (`SlicePickMany` with 0..100 threshold remains for legacy use)

### Time

Calendar-correct random datetimes and near-past/future helpers.

* `Datetime()` → between years 1..9999 (UTC), month/day validity handled
* `TimeInNearPast()` / `TimeInNearFuture()` → a few minutes around now

