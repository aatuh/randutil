//go:build randutil_policy
// +build randutil_policy

package randutil

import "github.com/aatuh/randutil/v2/core"

type disabledRoot struct{}

func (disabledRoot) Derive(_ string) (core.Source, error) {
	return nil, core.ErrDeterministicDisabled
}

func (disabledRoot) Close() error { return nil }

func newDeterministicRoot(_ []byte) Root {
	return disabledRoot{}
}
