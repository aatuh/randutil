# randutil

CSPRNG-first random utilities for Go. One small package that does the
common random things: bias-free integers, bytes, strings/tokens/emails,
shuffle/sample for slices, and realistic date/time helpers.

## Install

```bash
go get github.com/aatuh/randutil
```

## Quick start

```go
package main

import (
  "fmt"
  "github.com/aatuh/randutil"
)

func main() {
  // Numbers
  n := randutil.MustIntRange(10, 20)  // inclusive range
  f := randutil.MustFloat64()         // [0,1)
  b := randutil.MustBytes(16)         // 16 random bytes

  // Booleans
  ok := randutil.MustBool()

  // Strings / tokens / email
  s   := randutil.MustString(12)           // lower-case alnum
  hex := randutil.MustHex(32)              // 32 hex chars (16 bytes)
  b64 := randutil.MustBase64(24)           // encodes 24 random bytes
  tok := randutil.MustTokenURLSafe(24)     // URL-safe base64 (no padding)
  mail:= randutil.MustEmail(16)            // exact-length local@domain.com

  // Collections
  arr := []int{1,2,3,4,5}
  randutil.MustShuffle(arr)                // in-place Fisher–Yates
  top2 := randutil.MustSample(arr, 2)      // k without replacement
  pick := randutil.MustSlicePickOne(arr)
  pickedMany := randutil.MustSlicePickMany([]any{"a","b","c"}, 50)

  // Time
  t  := randutil.MustDatetime()            // valid calendar date/time
  p  := randutil.MustTimeInNearPast()
  fu := randutil.MustTimeInNearFuture()

  fmt.Println(n, f, len(b), ok, s, hex, b64, tok, mail, arr, top2, pick, pickedMany, t, p, fu)
}
```

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

Lower-case alnum by default; hex/base64 helpers; exact-length emails.

* `String(n)` and `GetWithCustomCharset(n, charset)`
* `Hex(len)` where `len` must be even (2 chars per byte)
* `Base64(byteLen)`, `TokenHex(byteLen)`, `TokenBase64(byteLen)`,
  `TokenURLSafe(byteLen)`
* `Email(totalLength)` builds `local@domain.com`, reserving 5 chars for
  "@" + ".com"; minimum length is 6.

### Collections (slices)

Unbiased in-place Fisher–Yates shuffle, sampling without replacement,
and simple picks.

* `Shuffle[T]([]T)`
* `Sample[T]([]T, k)` (returns a new slice of k unique items)
* `Perm(n)` → random permutation of 0..n-1
* `SlicePickOne[T]([]T)`
* `SlicePickMany([]any, chanceThreshold int)` draws a value 0..100 inclusive
  for each item and includes it when `value <= threshold`.

### Time

Calendar-correct random datetimes and near-past/future helpers.

* `Datetime()` → between years 1..9999 (UTC), month/day validity handled
* `TimeInNearPast()` / `TimeInNearFuture()` → a few minutes around now

## Related package: UUIDs

For UUID helpers (RFC 4122 v4, and more), see the companion module:

```go
// DEPRECATED: Use github.com/aatuh/randutil/uuid instead
import "github.com/aatuh/uuid"
```

Use `uuid.MustV4()` or `uuid.Parse(...)` as needed.

**Note: The uuid package is deprecated. Please use `github.com/aatuh/randutil/uuid` instead.**
