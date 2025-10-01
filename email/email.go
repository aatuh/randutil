package email

// Email returns a random email address with the specified options.
//
// Parameters:
//   - opts: EmailOptions configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
//   - error: An error if generation fails.
func Email(opts EmailOptions) (string, error) {
	return Default.Email(opts)
}

// MustEmail returns a random email address with the specified options.
// It panics if an error occurs.
//
// Parameters:
//   - opts: EmailOptions configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
func MustEmail(opts EmailOptions) string {
	result, err := Email(opts)
	if err != nil {
		panic(err)
	}
	return result
}

// EmailSimple returns a random email of exactly totalLength chars in the form
// local@domain.com (5 chars reserved for "@" + ".com"). This is the legacy
// behavior for backward compatibility.
//
// Parameters:
//   - totalLength: The exact total length of the email address.
//
// Returns:
//   - string: A random email address of the specified length.
//   - error: An error if totalLength is too small or if generation fails.
func EmailSimple(totalLength int) (string, error) {
	return Default.EmailSimple(totalLength)
}

// MustEmailSimple returns a random email of exactly totalLength chars.
// It panics if an error occurs.
//
// Parameters:
//   - totalLength: The exact total length of the email address.
//
// Returns:
//   - string: A random email address of the specified length.
func MustEmailSimple(totalLength int) string {
	result, err := EmailSimple(totalLength)
	if err != nil {
		panic(err)
	}
	return result
}

// EmailWithCustomLocal returns a random email with the specified local part.
//
// Parameters:
//   - localPart: The local part of the email address.
//
// Returns:
//   - string: A random email address with the specified local part.
//   - error: An error if generation fails.
func EmailWithCustomLocal(localPart string) (string, error) {
	return Email(EmailOptions{LocalPart: localPart})
}

// MustEmailWithCustomLocal returns a random email with the specified local part.
// It panics if an error occurs.
//
// Parameters:
//   - localPart: The local part of the email address.
//
// Returns:
//   - string: A random email address with the specified local part.
func MustEmailWithCustomLocal(localPart string) string {
	result, err := EmailWithCustomLocal(localPart)
	if err != nil {
		panic(err)
	}
	return result
}

// EmailWithCustomDomain returns a random email with the specified domain part.
//
// Parameters:
//   - domainPart: The domain part of the email address.
//
// Returns:
//   - string: A random email address with the specified domain part.
//   - error: An error if generation fails.
func EmailWithCustomDomain(domainPart string) (string, error) {
	return Email(EmailOptions{DomainPart: domainPart})
}

// MustEmailWithCustomDomain returns a random email with the specified domain part.
// It panics if an error occurs.
//
// Parameters:
//   - domainPart: The domain part of the email address.
//
// Returns:
//   - string: A random email address with the specified domain part.
func MustEmailWithCustomDomain(domainPart string) string {
	result, err := EmailWithCustomDomain(domainPart)
	if err != nil {
		panic(err)
	}
	return result
}

// EmailWithCustomTLD returns a random email with the specified TLD.
//
// Parameters:
//   - tld: The top-level domain (with or without leading dot).
//
// Returns:
//   - string: A random email address with the specified TLD.
//   - error: An error if generation fails.
func EmailWithCustomTLD(tld string) (string, error) {
	return Email(EmailOptions{TLD: tld})
}

// MustEmailWithCustomTLD returns a random email with the specified TLD.
// It panics if an error occurs.
//
// Parameters:
//   - tld: The top-level domain (with or without leading dot).
//
// Returns:
//   - string: A random email address with the specified TLD.
func MustEmailWithCustomTLD(tld string) string {
	result, err := EmailWithCustomTLD(tld)
	if err != nil {
		panic(err)
	}
	return result
}

// EmailWithRandomTLD returns a random email with a random TLD from commonTLDs.
//
// Returns:
//   - string: A random email address with a random TLD.
//   - error: An error if generation fails.
func EmailWithRandomTLD() (string, error) {
	return Email(EmailOptions{TLD: "random"})
}

// MustEmailWithRandomTLD returns a random email with a random TLD from commonTLDs.
// It panics if an error occurs.
//
// Returns:
//   - string: A random email address with a random TLD.
func MustEmailWithRandomTLD() string {
	result, err := EmailWithRandomTLD()
	if err != nil {
		panic(err)
	}
	return result
}

// EmailWithoutTLD returns a random email without a TLD.
//
// Returns:
//   - string: A random email address without a TLD.
//   - error: An error if generation fails.
func EmailWithoutTLD() (string, error) {
	return Email(EmailOptions{TLD: "none"})
}

// MustEmailWithoutTLD returns a random email without a TLD.
// It panics if an error occurs.
//
// Returns:
//   - string: A random email address without a TLD.
func MustEmailWithoutTLD() string {
	result, err := EmailWithoutTLD()
	if err != nil {
		panic(err)
	}
	return result
}
