package randutil

// Bool returns a secure random boolean.
//
// Returns:
//   - bool: A random boolean value.
//   - error
func Bool() (bool, error) {
	i, err := IntRange(0, 1)
	if err != nil {
		return false, err
	}
	return i == 1, nil
}

// Bool returns a secure random boolean. It panics if an error occurs.
//
// Returns:
//   - bool
func MustBool() bool {
	b, err := Bool()
	if err != nil {
		panic(err)
	}
	return b
}
