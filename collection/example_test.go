//go:build !randutil_policy
// +build !randutil_policy

package collection

import (
	"fmt"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func ExampleShuffle() {
	xs := []int{1, 2, 3}
	_ = Shuffle(xs)
	fmt.Println(len(xs))
	// Output: 3
}

func ExampleGenerator_WeightedSample() {
	src, err := adapters.DeterministicSource([]byte("seed"))
	if err != nil {
		panic(err)
	}
	gen := New[int](core.New(src))
	items := []int{1, 2, 3, 4}
	weights := []float64{1, 1, 1, 1}
	sample, _ := gen.WeightedSample(items, weights, 2)
	fmt.Println(len(sample))
	// Output: 2
}
