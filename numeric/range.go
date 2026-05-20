package numeric

// Constants to help define the full range for different integer types.
const (
	maxInt   = int(^uint(0) >> 1)
	minInt   = -maxInt - 1
	maxInt32 = int32(^uint32(0) >> 1) // 2147483647
	minInt32 = -maxInt32 - 1          // -2147483648
	maxInt64 = int64(^uint64(0) >> 1) //  9223372036854775807
	minInt64 = -maxInt64 - 1          // -9223372036854775808
)

// IntRange returns a secure random int in [minInclusive, maxInclusive].
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int: A random int in [minInclusive, maxInclusive].
//   - error: An error if crypto/rand fails.
func IntRange(minInclusive int, maxInclusive int) (int, error) {
	return Default().IntRange(minInclusive, maxInclusive)
}

// AnyInt returns a secure random int from the full range of int.
//
// Returns:
//   - int: A random int in the full int range.
//   - error: An error if crypto/rand fails.
func AnyInt() (int, error) {
	return IntRange(minInt, maxInt)
}

// Int32Range returns a secure random int32 in [minInclusive, maxInclusive].
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int32: A random int32 in [minInclusive, maxInclusive].
//   - error: An error if crypto/rand fails.
func Int32Range(minInclusive int32, maxInclusive int32) (int32, error) {
	return Default().Int32Range(minInclusive, maxInclusive)
}

// AnyInt32 returns a random int32 in the full int32 range.
//
// Returns:
//   - int32: A random int32 in the full int32 range.
//   - error: An error if crypto/rand fails.
func AnyInt32() (int32, error) {
	return Int32Range(minInt32, maxInt32)
}

// PositiveInt32 returns a secure random positive int32.
//
// Returns:
//   - int32: A random int32 in [1, 2147483647].
//   - error: An error if crypto/rand fails.
func PositiveInt32() (int32, error) {
	return Int32Range(1, maxInt32)
}

// NegativeInt32 returns a secure random negative int32.
//
// Returns:
//   - int32: A random int32 in [-2147483648, -1].
//   - error: An error if crypto/rand fails.
func NegativeInt32() (int32, error) {
	return Int32Range(minInt32, -1)
}

// Int64Range returns a secure random int64 in [minInclusive, maxInclusive].
//
// Parameters:
// - minInclusive: The minimum value (inclusive).
// - maxInclusive: The maximum value (inclusive).
//
// Returns:
//   - int64: A random int64 in [minInclusive, maxInclusive].
//   - error: An error if crypto/rand fails.
func Int64Range(minInclusive int64, maxInclusive int64) (int64, error) {
	return Default().Int64Range(minInclusive, maxInclusive)
}

// AnyInt64 returns a secure random int64 in the full int64 range.
//
// Returns:
//   - int64: A random int64 in [-9223372036854775808, 9223372036854775807].
//   - error: An error if crypto/rand fails.
func AnyInt64() (int64, error) {
	return Int64Range(minInt64, maxInt64)
}

// PositiveInt64 returns a secure random positive int64.
//
// Returns:
//   - int64: A random int64 in [1, 9223372036854775807].
//   - error: An error if crypto/rand fails.
func PositiveInt64() (int64, error) {
	return Int64Range(1, maxInt64)
}

// NegativeInt64 returns a secure random negative int64.
//
// Returns:
//   - int64: A random int64 in [-9223372036854775808, -1].
//   - error: An error if crypto/rand fails.
func NegativeInt64() (int64, error) {
	return Int64Range(minInt64, -1)
}
