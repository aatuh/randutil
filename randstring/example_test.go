package randstring

import "fmt"

func ExampleString() {
	s, _ := String(8)
	fmt.Println(len(s))
	// Output: 8
}
