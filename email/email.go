package email

// Email returns a random email address with the specified options.
//
// Parameters:
//   - opts: Options configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
//   - error: An error if generation fails.
func Email(opts Options) (string, error) {
	return Default().Email(opts)
}

// Simple returns a random email of exactly totalLength chars in the form
// local@domain.com (5 chars reserved for "@" + ".com"). This is the legacy
// behavior for backward compatibility.
//
// Parameters:
//   - totalLength: The exact total length of the email address.
//
// Returns:
//   - string: A random email address of the specified length.
//   - error: An error if totalLength is too small or if generation fails.
func Simple(totalLength int) (string, error) {
	return Default().Simple(totalLength)
}

// WithCustomLocal returns a random email with the specified local part.
//
// Parameters:
//   - localPart: The local part of the email address.
//
// Returns:
//   - string: A random email address with the specified local part.
//   - error: An error if generation fails.
func WithCustomLocal(localPart string) (string, error) {
	return Email(Options{LocalPart: localPart})
}

// WithCustomDomain returns a random email with the specified domain part.
//
// Parameters:
//   - domainPart: The domain part of the email address.
//
// Returns:
//   - string: A random email address with the specified domain part.
//   - error: An error if generation fails.
func WithCustomDomain(domainPart string) (string, error) {
	return Email(Options{DomainPart: domainPart})
}

// WithCustomTLD returns a random email with the specified TLD.
//
// Parameters:
//   - tld: The top-level domain (with or without leading dot).
//
// Returns:
//   - string: A random email address with the specified TLD.
//   - error: An error if generation fails.
func WithCustomTLD(tld string) (string, error) {
	return Email(Options{TLD: tld})
}

// WithRandomTLD returns a random email with a random TLD from commonTLDs.
//
// Returns:
//   - string: A random email address with a random TLD.
//   - error: An error if generation fails.
func WithRandomTLD() (string, error) {
	return Email(Options{TLD: "random"})
}

// WithoutTLD returns a random email without a TLD.
//
// Returns:
//   - string: A random email address without a TLD.
//   - error: An error if generation fails.
func WithoutTLD() (string, error) {
	return Email(Options{TLD: "none"})
}
