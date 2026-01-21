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

// MustEmail returns a random email address with the specified options.
// It panics if an error occurs.
//
// Parameters:
//   - opts: Options configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
func MustEmail(opts Options) string {
	result, err := Email(opts)
	if err != nil {
		panic(err)
	}
	return result
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

// MustSimple returns a random email of exactly totalLength chars.
// It panics if an error occurs.
//
// Parameters:
//   - totalLength: The exact total length of the email address.
//
// Returns:
//   - string: A random email address of the specified length.
func MustSimple(totalLength int) string {
	result, err := Simple(totalLength)
	if err != nil {
		panic(err)
	}
	return result
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

// MustWithCustomLocal returns a random email with the specified local part.
// It panics if an error occurs.
//
// Parameters:
//   - localPart: The local part of the email address.
//
// Returns:
//   - string: A random email address with the specified local part.
func MustWithCustomLocal(localPart string) string {
	result, err := WithCustomLocal(localPart)
	if err != nil {
		panic(err)
	}
	return result
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

// MustWithCustomDomain returns a random email with the specified domain part.
// It panics if an error occurs.
//
// Parameters:
//   - domainPart: The domain part of the email address.
//
// Returns:
//   - string: A random email address with the specified domain part.
func MustWithCustomDomain(domainPart string) string {
	result, err := WithCustomDomain(domainPart)
	if err != nil {
		panic(err)
	}
	return result
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

// MustWithCustomTLD returns a random email with the specified TLD.
// It panics if an error occurs.
//
// Parameters:
//   - tld: The top-level domain (with or without leading dot).
//
// Returns:
//   - string: A random email address with the specified TLD.
func MustWithCustomTLD(tld string) string {
	result, err := WithCustomTLD(tld)
	if err != nil {
		panic(err)
	}
	return result
}

// WithRandomTLD returns a random email with a random TLD from commonTLDs.
//
// Returns:
//   - string: A random email address with a random TLD.
//   - error: An error if generation fails.
func WithRandomTLD() (string, error) {
	return Email(Options{TLD: "random"})
}

// MustWithRandomTLD returns a random email with a random TLD from commonTLDs.
// It panics if an error occurs.
//
// Returns:
//   - string: A random email address with a random TLD.
func MustWithRandomTLD() string {
	result, err := WithRandomTLD()
	if err != nil {
		panic(err)
	}
	return result
}

// WithoutTLD returns a random email without a TLD.
//
// Returns:
//   - string: A random email address without a TLD.
//   - error: An error if generation fails.
func WithoutTLD() (string, error) {
	return Email(Options{TLD: "none"})
}

// MustWithoutTLD returns a random email without a TLD.
// It panics if an error occurs.
//
// Returns:
//   - string: A random email address without a TLD.
func MustWithoutTLD() string {
	result, err := WithoutTLD()
	if err != nil {
		panic(err)
	}
	return result
}
