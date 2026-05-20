//go:build randutil_policy
// +build randutil_policy

package adapters

import "github.com/aatuh/randutil/v2/core"

// DeterministicSource returns an error when policy mode is enabled.
func DeterministicSource(_ []byte) (core.Source, error) {
	return nil, core.ErrDeterministicDisabled
}

// DeterministicSourceWithLabel returns an error when policy mode is enabled.
func DeterministicSourceWithLabel(_ []byte, _ string) (core.Source, error) {
	return nil, core.ErrDeterministicDisabled
}
