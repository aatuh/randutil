//go:build !randutil_policy
// +build !randutil_policy

package dist

import (
	"fmt"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func ExampleGenerator() {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		panic(err)
	}
	gen := New(core.New(src))
	_, _ = gen.Normal(0, 1)
}

func ExampleGenerator_Bernoulli() {
	src := testutil.NewSeqReader(testutil.Float64Bytes(0.25))
	gen := New(core.New(src))
	v, _ := gen.Bernoulli(0.5)
	fmt.Println(v)
	// Output: true
}
