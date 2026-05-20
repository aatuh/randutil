//go:build !randutil_policy
// +build !randutil_policy

package core_test

import (
	"fmt"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func ExampleNew() {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		panic(err)
	}
	rng := core.New(src)
	b, _ := rng.Bytes(4)
	fmt.Println(len(b))
	// Output: 4
}
