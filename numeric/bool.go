package numeric

// Bool returns a secure random boolean.
//
// Returns:
//   - bool: A random boolean value.
//   - error
func Bool() (bool, error) {
	return def.Bool()
}

// MustBool returns a secure random boolean. It panics if an error occurs.
//
// Returns:
//   - bool: A random boolean value.
func MustBool() bool {
	b, err := Bool()
	if err != nil {
		panic(err)
	}
	return b
}
