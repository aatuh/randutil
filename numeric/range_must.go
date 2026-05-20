//go:build randutil_must
// +build randutil_must

package numeric

// MustIntRange returns a secure random int in [minInclusive, maxInclusive].
// It panics if an error occurs.
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int: A random int in [minInclusive, maxInclusive].
func MustIntRange(minInclusive int, maxInclusive int) int {
	i, err := IntRange(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return i
}

// MustAnyInt returns a secure random int from the full range of int.
// It panics if an error occurs.
//
// Returns:
//   - int: A random int in the full int range.
func MustAnyInt() int {
	return MustIntRange(minInt, maxInt)
}

// MustInt32Range returns a secure random int32 in [minInclusive, maxInclusive].
// It panics if an error occurs.
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int32: A random int32 in [minInclusive, maxInclusive].
func MustInt32Range(minInclusive int32, maxInclusive int32) int32 {
	i, err := Int32Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return i
}

// MustAnyInt32 returns a random int32 in the full int32 range.
// It panics if an error occurs.
//
// Returns:
//   - int32: A random int32 in the full int32 range.
func MustAnyInt32() int32 {
	return MustInt32Range(minInt32, maxInt32)
}

// MustPositiveInt32 returns a secure random positive int32.
// It panics if an error occurs.
//
// Returns:
//   - int32: A random int32 in [1, 2147483648].
func MustPositiveInt32() int32 {
	return MustInt32Range(1, maxInt32)
}

// MustNegativeInt32 returns a secure random negative int32.
// It panics if an error occurs.
//
// Returns:
//   - int32: A random int32 in [-2147483647, -1].
func MustNegativeInt32() int32 {
	return MustInt32Range(minInt32, -1)
}

// MustInt64Range returns a secure random int64 in [minInclusive, maxInclusive].
// It panics if an error occurs.
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int64: A random int64 in [minInclusive, maxInclusive].
func MustInt64Range(minInclusive int64, maxInclusive int64) int64 {
	i, err := Int64Range(minInclusive, maxInclusive)
	if err != nil {
		panic(err)
	}
	return i
}

// MustAnyInt64 returns a secure random int64 in the full int64 range.
// It panics if an error occurs.
//
// Returns:
//   - int64: A random int64 in [-9223372036854775808, 9223372036854775807].
func MustAnyInt64() int64 {
	return MustInt64Range(minInt64, maxInt64)
}

// MustPositiveInt64 returns a secure random positive int64.
// It panics if an error occurs.
//
// Returns:
//   - int64: A random int64 in [1, 9223372036854775807].
func MustPositiveInt64() int64 {
	return MustInt64Range(1, maxInt64)
}

// MustNegativeInt64 returns a secure random negative int64.
// It panics if an error occurs.
//
// Returns:
//   - int64: A random int64 in [-9223372036854775808, -1].
func MustNegativeInt64() int64 {
	return MustInt64Range(minInt64, -1)
}
