//go:build !randutil_policy
// +build !randutil_policy

package uuid

import (
	"fmt"
	"time"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func ExampleV4() {
	u, _ := V4()
	fmt.Println(len(u) == 36)
	// Output: true
}

func ExampleGenerator_V7() {
	fixed := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	src, err := adapters.DeterministicSourceWithLabel([]byte("seed"), "uuid")
	if err != nil {
		panic(err)
	}
	gen := NewWithClock(
		core.New(src),
		func() time.Time { return fixed },
	)
	u, _ := gen.V7()
	fmt.Println(u)
	// Output: 018cc820-d888-7a3b-ae5c-d5bc8d457263
}
