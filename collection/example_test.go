package collection

import "fmt"

func ExampleShuffle() {
	xs := []int{1, 2, 3}
	_ = Shuffle(xs)
	fmt.Println(len(xs))
	// Output: 3
}
