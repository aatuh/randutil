//go:build !randutil_policy
// +build !randutil_policy

package randutil

func newDeterministicRoot(seed []byte) Root {
	return newSeedRoot(seed)
}
