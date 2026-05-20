package nanoid

// ID returns a NanoID with the default length.
func ID() (string, error) {
	return Default().ID(DefaultLength)
}

// IDWithLength returns a NanoID with the requested length.
func IDWithLength(length int) (string, error) {
	return Default().ID(length)
}
