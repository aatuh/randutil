//go:build randutil_must
// +build randutil_must

package email

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
