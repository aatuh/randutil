package numeric

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}
