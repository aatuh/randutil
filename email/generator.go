package email

import (
	"fmt"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/randstring"
)

// Generator builds email-related random operations using a core RNG.
// It also maintains a randstring.Generator that shares the same entropy.
//
// Concurrency: safe for concurrent use if the underlying RNG is safe.
type Generator struct {
	rng     core.RNG
	strings *randstring.Generator
}

// New returns an email Generator. If rng is nil, crypto/rand is used.
func New(rng core.RNG) *Generator {
	if rng == nil {
		rng = core.New(nil)
	}
	return &Generator{
		rng:     rng,
		strings: randstring.New(rng),
	}
}

// NewWithSource returns an email Generator bound to src.
func NewWithSource(src core.Source) *Generator {
	return New(core.New(src))
}

var defaultGenerator = New(nil)

// Default returns the package-wide default generator.
func Default() *Generator {
	return defaultGenerator
}

// Options configures email generation behavior.
type Options struct {
	// LocalPart specifies the local part of the email. If empty, a random
	// string will be generated. If set, this exact value will be used.
	LocalPart string

	// DomainPart specifies the domain part of the email. If empty, a random
	// string will be generated. If set, this exact value will be used.
	DomainPart string

	// TLD specifies the top-level domain. If empty, ".com" will be used.
	// If set to "random", a random TLD from commonTLDs will be used.
	// If set to any other value, that exact TLD will be used.
	// If set to "none", no TLD will be added.
	TLD string

	// TotalLength specifies the exact total length of the email address.
	// If 0, the length will be calculated automatically based on the parts.
	// This option is ignored if LocalPart, DomainPart, or TLD are specified.
	TotalLength int
}

// Email returns a random email address with the specified options.
//
// Parameters:
//   - opts: Options configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
//   - error: An error if generation fails.
func (g *Generator) Email(opts Options) (string, error) {
	if opts.TotalLength > 0 && opts.LocalPart == "" &&
		opts.DomainPart == "" && opts.TLD == "" {
		return g.Simple(opts.TotalLength)
	}

	tld := ""
	switch opts.TLD {
	case "":
		tld = ".com"
	case "random":
		idx, err := g.rng.Uint64n(uint64(len(commonTLDs)))
		if err != nil {
			return "", err
		}
		tld = "." + commonTLDs[idx]
	case "none":
		tld = ""
	default:
		if opts.TLD[0] != '.' {
			tld = "." + opts.TLD
		} else {
			tld = opts.TLD
		}
	}

	local := opts.LocalPart
	if local == "" {
		var err error
		local, err = g.strings.String(5)
		if err != nil {
			return "", err
		}
	}

	domain := opts.DomainPart
	if domain == "" {
		var err error
		domain, err = g.strings.String(5)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s@%s%s", local, domain, tld), nil
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
func (g *Generator) Simple(totalLength int) (string, error) {
	if totalLength < 7 {
		return "", ErrTotalLengthTooSmall
	}
	body := totalLength - 5
	localLen := body / 2
	if localLen == 0 {
		localLen = 1
	}
	domainLen := body - localLen
	if domainLen == 0 {
		domainLen = 1
		localLen = body - domainLen
	}
	local, err := g.strings.String(localLen)
	if err != nil {
		return "", err
	}
	domain, err := g.strings.String(domainLen)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s@%s.com", local, domain), nil
}

// Common TLDs for random selection.
var commonTLDs = []string{
	"com", "org", "net", "edu", "gov", "mil", "int", "co", "io", "ai",
	"app", "dev", "tech", "online", "site", "store", "blog", "news",
	"info", "biz", "name", "me", "us", "uk", "ca", "de", "fr", "jp",
}
