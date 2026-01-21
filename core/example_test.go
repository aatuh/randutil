package core_test

import (
	"fmt"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func ExampleNew() {
	rng := core.New(adapters.DeterministicSource([]byte("seed")))
	b, _ := rng.Bytes(4)
	fmt.Println(len(b))
	// Output: 4
}
