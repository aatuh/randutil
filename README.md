# randutil

Secure random helpers for numbers, strings/bytes, time, and collections.

## Install

```go
import "github.com/aatuh/random"
```

## Quick start

```go
s := random.MustString(12)           // a1b2c3...
n := random.MustIntRange(1, 10)      // 1..10 inclusive
b := random.MustBool()               // true/false
t := random.MustDatetime()           // random time
e := random.MustEmail(16)            // local@domain.com (~16 chars)
c := random.MustChoice("a", "b", "c") // one of the args
```

## Numbers

- Int ranges: `IntRange`, `AnyInt`
- int32: `Int32Range`, `AnyInt32`, `PositiveInt32`, `NegativeInt32`
- int64: `Int64Range`, `AnyInt64`, `PositiveInt64`, `NegativeInt64`

```go
i, err := random.IntRange(10, 20)
```

## Strings and bytes

- Random strings: `String(length)` (lowercase+digits)
- Custom charset: `GetWithCustomCharset(length, charset)`
- Base64 bytes: `Base64(byteLen)`
- Hex string: `Hex(strLen)` (strLen must be even)
- Slices: `StringSlice(count, minLen, maxLen)`

```go
hex, _ := random.Hex(32)
b64, _ := random.Base64(24)
word, _ := random.String(8)
```

## Collections

- Pick/shuffle: `SlicePickOne`, `SlicePickMany`, `Shuffle`, `Choice`

```go
v, _ := random.SlicePickOne([]any{"x", 2, true})
items := random.MustSlicePickMany([]any{1,2,3,4}, 50) // ~50% picked
arr := []any{1,2,3,4}; random.MustShuffle(arr)
```

## Time

- `Datetime()` random `time.Time`
- `TimeInNearPast()` / `TimeInNearFuture()` (±5–10 minutes)

```go
past := random.MustTimeInNearPast()
future := random.MustTimeInNearFuture()
```

## Email

- `Email(totalLength)` builds `local@domain.com` of approx total length

```go
addr := random.MustEmail(20)
```

## Notes

- Uses `crypto/rand` throughout.
- All numeric ranges are inclusive.
- `Must*` variants panic on error.
