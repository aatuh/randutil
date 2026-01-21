package numeric

import "fmt"

func ExampleIntRange() {
	v, _ := IntRange(1, 10)
	fmt.Println(v >= 1 && v <= 10)
	// Output: true
}
