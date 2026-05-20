//go:build !randutil_policy
// +build !randutil_policy

package adapters

import (
	"bytes"
	"fmt"
)

func ExampleDeterministicSource() {
	seed := []byte("seed")
	a, err := DeterministicSource(seed)
	if err != nil {
		panic(err)
	}
	b, err := DeterministicSource(seed)
	if err != nil {
		panic(err)
	}
	bufA := make([]byte, 4)
	bufB := make([]byte, 4)
	_, _ = a.Read(bufA)
	_, _ = b.Read(bufB)
	fmt.Println(bytes.Equal(bufA, bufB))
	// Output: true
}
