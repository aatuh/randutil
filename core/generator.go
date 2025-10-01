package core

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"io"
	"math/big"
)

// Generator builds numbers and bytes using an entropy source.
// Zero-value uses the current package source via Reader().
type Generator struct {
	R io.Reader
}

// New returns a core Generator. If src is nil, the package default is used.
//
// Parameters:
//   - src: The entropy source to use.
//
// Returns:
//   - *Generator: A new core Generator.
func New(src io.Reader) *Generator {
	if src == nil {
		return &Generator{R: nil} // Will use Reader() via reader() method
	}
	return &Generator{R: src}
}

func (g *Generator) reader() io.Reader {
	if g == nil || g.R == nil {
		return Reader()
	}
	return g.R
}

// Read implements io.Reader by delegating to the generator's source.
// This allows core.Generator to be used anywhere an entropy io.Reader is
// required.
//
// Parameters:
//   - p: The byte slice to read into.
//
// Returns:
//   - int: The number of bytes read.
//   - error: An error if the source fails.
func (g *Generator) Read(p []byte) (int, error) {
	return g.reader().Read(p)
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
		return nil, ErrInvalidN
	}
	if n == 0 {
		return []byte{}, nil
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
	_, err := io.ReadFull(g.reader(), b)
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
		return 0, ErrInvalidRange
	}
	var (
		max   = ^uint64(0)
		limit = max - (max % n)
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
		return 0, ErrInvalidN
	}
	u, err := g.Uint64n(uint64(n))
	if err != nil {
		return 0, err
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
		return 0, ErrInvalidN
	}
	u, err := g.Uint64n(uint64(n))
	if err != nil {
		return 0, err
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
// This is optimized to avoid the big.Int path used by IntRange.
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
	// Use big.Int math to avoid overflow on 64-bit platforms.
	bigMin := big.NewInt(int64(minInclusive))
	bigMax := big.NewInt(int64(maxInclusive))
	diff := new(big.Int).Sub(bigMax, bigMin)
	diff.Add(diff, big.NewInt(1))
	if diff.Sign() <= 0 {
		return 0, ErrInvalidRangeNonPositive
	}
	n, err := g.bigInt(diff) // [0, diff)
	if err != nil {
		return 0, err
	}
	res := new(big.Int).Add(n, bigMin)
	// Final bounds check. Ensures it fits in int64 -> int.
	if res.BitLen() > 63 {
		return 0, ErrResultOutOfRange
	}
	return int(res.Int64()), nil
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
	rng := big.NewInt(diff)
	n, err := g.bigInt(rng)
	if err != nil {
		return 0, err
	}
	return int32(n.Int64()) + minInclusive, nil
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

	// Convert the bounds to big.Int.
	bigMin := big.NewInt(minInclusive)
	bigMax := big.NewInt(maxInclusive)

	// diff = (max - min + 1) as a big.Int
	bigDiff := new(big.Int).Sub(bigMax, bigMin)
	bigDiff.Add(bigDiff, big.NewInt(1))

	if bigDiff.Sign() <= 0 {
		return 0, ErrInvalidRangeNonPositive
	}

	// Generate a random big.Int in [0, bigDiff).
	n, err := g.bigInt(bigDiff)
	if err != nil {
		return 0, err
	}

	// Shift by minInclusive to get [minInclusive, maxInclusive].
	bigResult := new(big.Int).Add(n, bigMin)

	// Final check: ensure it's in int64 range (it should be).
	if bigResult.BitLen() > 63 {
		return 0, ErrResultOutOfRange
	}

	return bigResult.Int64(), nil
}

// bigInt returns a random big.Int in [0, max) using the generator's source.
func (g *Generator) bigInt(max *big.Int) (*big.Int, error) {
	reader := g.reader()
	if reader == nil {
		return nil, errors.New("no entropy source available")
	}
	return rand.Int(reader, max)
}
