package email

import (
	"errors"
	"fmt"
	"io"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/randstring"
)

// Generator builds email-related random operations using a core generator.
// It also maintains a randstring.Generator that shares the same entropy.
type Generator struct {
	G  core.Generator
	RS randstring.Generator
}

// New returns an email Generator. If src is nil, the core default is used.
func New(src io.Reader) *Generator {
	g := core.Generator{R: src}
	return &Generator{
		G:  g,
		RS: randstring.Generator{G: g},
	}
}

// Default is the package-wide default generator.
var Default = New(nil)

// EmailOptions configures email generation behavior.
type EmailOptions struct {
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
//   - opts: EmailOptions configuring the email generation behavior.
//
// Returns:
//   - string: A random email address.
//   - error: An error if generation fails.
func (g *Generator) Email(opts EmailOptions) (string, error) {
	// Legacy: only TotalLength provided -> exact-length local@domain.com.
	if opts.TotalLength > 0 && opts.LocalPart == "" &&
		opts.DomainPart == "" && opts.TLD == "" {
		return g.EmailSimple(opts.TotalLength)
	}

	// Determine TLD first.
	tld := ""
	switch opts.TLD {
	case "":
		tld = ".com"
	case "random":
		idx, err := g.G.Uint64n(uint64(len(commonTLDs)))
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

	// If any override is present, ignore TotalLength entirely.
	useTotal := (opts.LocalPart == "" && opts.DomainPart == "" && opts.TLD == "" &&
		opts.TotalLength > 0)

	var local, domain string
	var err error

	if opts.LocalPart != "" {
		local = opts.LocalPart
	} else if useTotal {
		// Recompute exact lengths for default ".com".
		available := opts.TotalLength - len(".com") - 1 // "@"
		if available < 2 {
			return "", errors.New("totalLength too small for email parts")
		}
		localLen := available / 2
		if localLen == 0 {
			localLen = 1
		}
		domainLen := available - localLen
		if domainLen == 0 {
			domainLen = 1
			localLen = available - domainLen
		}
		local, err = g.RS.String(localLen)
		if err != nil {
			return "", err
		}
		domain, err = g.RS.String(domainLen)
		if err != nil {
			return "", err
		}
		// Done generating both parts with exact length.
		return fmt.Sprintf("%s@%s%s", local, domain, tld), nil
	} else {
		// No override, no TotalLength constraint -> default sized random local.
		local, err = g.RS.String(5)
		if err != nil {
			return "", err
		}
	}

	if opts.DomainPart != "" {
		domain = opts.DomainPart
	} else if domain == "" {
		// Default sized random domain when not set explicitly.
		domain, err = g.RS.String(5)
		if err != nil {
			return "", err
		}
	}

	return fmt.Sprintf("%s@%s%s", local, domain, tld), nil
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
func (g *Generator) EmailSimple(totalLength int) (string, error) {
	if totalLength < 7 {
		return "", errors.New("totalLength must be at least 7")
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
	local, err := g.RS.String(localLen)
	if err != nil {
		return "", err
	}
	domain, err := g.RS.String(domainLen)
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
