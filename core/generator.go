package core

import (
	crand "crypto/rand"
	"encoding/binary"
	"io"
	"math/big"
)

const (
	maxInt   = int(^uint(0) >> 1)
	minInt   = -maxInt - 1
	maxInt64 = int64(^uint64(0) >> 1)
	maxInt32 = int64(^uint32(0) >> 1)
	minInt32 = -maxInt32 - 1
)

// Generator builds numbers and bytes using an entropy source.
// Zero-value uses crypto/rand.Reader.
//
// Concurrency: safe for concurrent use if the underlying Source is safe.
type Generator struct {
	src Source
}

// New returns a core Generator. If src is nil, crypto/rand.Reader is used.
//
// Parameters:
//   - src: The entropy source to use.
//
// Returns:
//   - *Generator: A new core Generator.
func New(src Source) *Generator {
	if src == nil {
		src = crand.Reader
	}
	return &Generator{src: src}
}

func (g *Generator) source() Source {
	if g == nil || g.src == nil {
		return crand.Reader
	}
	return g.src
}

// Source returns the underlying entropy source (or crypto/rand.Reader).
//
// Returns:
//   - Source: The configured entropy source.
func (g *Generator) Source() Source {
	return g.source()
}

// Read implements io.Reader by delegating to the generator's source.
//
// Parameters:
//   - p: The byte slice to read into.
//
// Returns:
//   - int: The number of bytes read.
//   - error: An error if the source fails.
func (g *Generator) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	return io.ReadFull(g.source(), p)
}

// Bytes returns n random bytes from the generator's entropy source.
//
// Parameters:
//   - n: The number of random bytes to generate.
//
// Returns:
//   - []byte: n random bytes from the generator's entropy source.
//   - error: An error if n < 0 or if entropy fails.
func (g *Generator) Bytes(n int) ([]byte, error) {
	if n < 0 {
		return nil, ErrNegativeLength
	}
	buf := make([]byte, n)
	if err := g.Fill(buf); err != nil {
		return nil, err
	}
	return buf, nil
}

// Fill populates b with random bytes from the generator's entropy source.
//
// Parameters:
//   - b: The byte slice to fill with random data.
//
// Returns:
//   - error: An error if entropy fails.
func (g *Generator) Fill(b []byte) error {
	if len(b) == 0 {
		return nil
	}
	_, err := io.ReadFull(g.source(), b)
	if err != nil {
		for i := range b {
			b[i] = 0
		}
	}
	return err
}

// Uint64 returns a random uint64 from the generator's entropy source.
//
// Returns:
//   - uint64: A random uint64 value.
//   - error: An error if entropy fails.
func (g *Generator) Uint64() (uint64, error) {
	var b [8]byte
	if err := g.Fill(b[:]); err != nil {
		return 0, err
	}
	return binary.LittleEndian.Uint64(b[:]), nil
}

// Uint64n returns a uniform random integer in [0, n) using rejection
// sampling to avoid modulo bias. n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - uint64: A random uint64 in [0, n).
//   - error: An error if n == 0 or if entropy fails.
func (g *Generator) Uint64n(n uint64) (uint64, error) {
	if n == 0 {
		return 0, ErrNonPositiveBound
	}
	var (
		maxUint = ^uint64(0)
		limit   = maxUint - (maxUint % n)
	)
	for {
		u, err := g.Uint64()
		if err != nil {
			return 0, err
		}
		if u < limit {
			return u % n, nil
		}
	}
}

// Intn returns a uniform random int in [0, n). n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - int: A random int in [0, n).
//   - error: An error if n <= 0 or if entropy fails.
func (g *Generator) Intn(n int) (int, error) {
	if n <= 0 {
		return 0, ErrNonPositiveBound
	}
	u, err := g.Uint64n(uint64(n))
	if err != nil {
		return 0, err
	}
	if u > uint64(maxInt) {
		return 0, ErrResultOutOfRange
	}
	return int(u), nil
}

// Int64n returns a uniform random int64 in [0, n). n must be > 0.
//
// Parameters:
//   - n: The upper bound (exclusive).
//
// Returns:
//   - int64: A random int64 in [0, n).
//   - error: An error if n <= 0 or if entropy fails.
func (g *Generator) Int64n(n int64) (int64, error) {
	if n <= 0 {
		return 0, ErrNonPositiveBound
	}
	u, err := g.Uint64n(uint64(n))
	if err != nil {
		return 0, err
	}
	if u > uint64(maxInt64) {
		return 0, ErrResultOutOfRange
	}
	return int64(u), nil
}

// Float64 returns a uniform random float64 in [0.0, 1.0) with 53 bits
// of precision built from the generator's entropy source.
//
// Returns:
//   - float64: A random float64 in [0.0, 1.0).
//   - error: An error if entropy fails.
func (g *Generator) Float64() (float64, error) {
	var b [8]byte
	if err := g.Fill(b[:]); err != nil {
		return 0, err
	}
	u := binary.LittleEndian.Uint64(b[:]) >> 11
	const denom = 1 << 53
	return float64(u) / float64(denom), nil
}

// Bool returns a random boolean from the generator's entropy source.
//
// Returns:
//   - bool: A random boolean value.
//   - error: An error if entropy fails.
func (g *Generator) Bool() (bool, error) {
	u, err := g.Uint64()
	if err != nil {
		return false, err
	}
	return (u & 1) == 1, nil
}

// IntRange returns a secure random int in [minInclusive, maxInclusive].
//
// Parameters:
//   - minInclusive: The minimum value (inclusive).
//   - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int: A random int in [minInclusive, maxInclusive].
//   - error: An error if minInclusive > maxInclusive or if entropy fails.
func (g *Generator) IntRange(minInclusive int, maxInclusive int) (int, error) {
	if minInclusive > maxInclusive {
		return 0, ErrMinGreaterThanMax
	}
	min64 := int64(minInclusive)
	max64 := int64(maxInclusive)
	if span, ok := spanInt64(min64, max64); ok && span > 0 &&
		span <= uint64(maxInt64)+1 {
		u64, err := g.Uint64n(span)
		if err != nil {
			return 0, err
		}
		u, err := uint64ToInt64(u64)
		if err != nil {
			return 0, err
		}
		return int64ToInt(min64 + u)
	}

	bigMin := big.NewInt(min64)
	bigMax := big.NewInt(max64)
	diff := new(big.Int).Sub(bigMax, bigMin)
	diff.Add(diff, big.NewInt(1))
	if diff.Sign() <= 0 {
		return 0, ErrInvalidRangeNonPositive
	}
	n, err := g.bigInt(diff)
	if err != nil {
		return 0, err
	}
	res := new(big.Int).Add(n, bigMin)
	if res.BitLen() > 63 {
		return 0, ErrResultOutOfRange
	}
	return int64ToInt(res.Int64())
}

// Int32Range returns a secure random int32 in [minInclusive, maxInclusive].
//
// Parameters:
//   - minInclusive: The minimum value (inclusive).
//   - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int32: A random int32 in [minInclusive, maxInclusive].
//   - error: An error if minInclusive > maxInclusive or if entropy fails.
func (g *Generator) Int32Range(minInclusive int32, maxInclusive int32) (int32, error) {
	if minInclusive > maxInclusive {
		return 0, ErrMinGreaterThanMax
	}
	diff := int64(maxInclusive) - int64(minInclusive) + 1
	if diff <= 0 {
		return 0, ErrInvalidRangeNonPositive
	}
	u64, err := g.Uint64n(uint64(diff))
	if err != nil {
		return 0, err
	}
	u, err := uint64ToInt64(u64)
	if err != nil {
		return 0, err
	}
	return int64ToInt32(int64(minInclusive) + u)
}

// Int64Range returns a secure random int64 in [minInclusive, maxInclusive].
//
// Parameters:
//   - minInclusive: The minimum value (inclusive).
//   - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int64: A random int64 in [minInclusive, maxInclusive].
//   - error: An error if minInclusive > maxInclusive or if entropy fails.
func (g *Generator) Int64Range(minInclusive int64, maxInclusive int64) (int64, error) {
	if minInclusive > maxInclusive {
		return 0, ErrMinGreaterThanMax
	}
	if span, ok := spanInt64(minInclusive, maxInclusive); ok && span > 0 &&
		span <= uint64(maxInt64)+1 {
		u64, err := g.Uint64n(span)
		if err != nil {
			return 0, err
		}
		u, err := uint64ToInt64(u64)
		if err != nil {
			return 0, err
		}
		res := minInclusive + u
		if res < minInclusive || res > maxInclusive {
			return 0, ErrResultOutOfRange
		}
		return res, nil
	}

	bigMin := big.NewInt(minInclusive)
	bigMax := big.NewInt(maxInclusive)
	bigDiff := new(big.Int).Sub(bigMax, bigMin)
	bigDiff.Add(bigDiff, big.NewInt(1))
	if bigDiff.Sign() <= 0 {
		return 0, ErrInvalidRangeNonPositive
	}
	n, err := g.bigInt(bigDiff)
	if err != nil {
		return 0, err
	}
	bigResult := new(big.Int).Add(n, bigMin)
	if bigResult.BitLen() > 63 {
		return 0, ErrResultOutOfRange
	}
	return bigResult.Int64(), nil
}

// bigInt returns a random big.Int in [0, max) using the generator's source.
func (g *Generator) bigInt(upper *big.Int) (*big.Int, error) {
	return crand.Int(g.source(), upper)
}

func spanInt64(minInclusive int64, maxInclusive int64) (uint64, bool) {
	if minInclusive > maxInclusive {
		return 0, false
	}
	if minInclusive >= 0 {
		diff := maxInclusive - minInclusive
		if diff < 0 {
			return 0, false
		}
		span, err := int64ToUint64(diff)
		if err != nil {
			return 0, false
		}
		return span + 1, true
	}
	if maxInclusive < 0 {
		diff := maxInclusive - minInclusive
		if diff < 0 {
			return 0, false
		}
		span, err := int64ToUint64(diff)
		if err != nil {
			return 0, false
		}
		return span + 1, true
	}
	absMin := absInt64ToUint64(minInclusive)
	maxU, err := int64ToUint64(maxInclusive)
	if err != nil {
		return 0, false
	}
	if absMin > ^uint64(0)-maxU-1 {
		return 0, false
	}
	return absMin + maxU + 1, true
}

func absInt64ToUint64(v int64) uint64 {
	if v >= 0 {
		// #nosec G115 -- v is non-negative by construction.
		return uint64(v)
	}
	// #nosec G115 -- two's complement absolute value for negative int64.
	return uint64(^v) + 1
}

func uint64ToInt64(n uint64) (int64, error) {
	if n > uint64(maxInt64) {
		return 0, ErrResultOutOfRange
	}
	return int64(n), nil
}

func int64ToUint64(n int64) (uint64, error) {
	if n < 0 {
		return 0, ErrResultOutOfRange
	}
	// #nosec G115 -- conversion guarded by range check.
	return uint64(n), nil
}

func int64ToInt(n int64) (int, error) {
	if n < int64(minInt) || n > int64(maxInt) {
		return 0, ErrResultOutOfRange
	}
	return int(n), nil
}

func int64ToInt32(n int64) (int32, error) {
	if n < minInt32 || n > maxInt32 {
		return 0, ErrResultOutOfRange
	}
	return int32(n), nil
}
