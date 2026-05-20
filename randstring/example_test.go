package randstring

import "fmt"

func ExampleString() {
	s, _ := String(8)
	fmt.Println(len(s))
	// Output: 8
}

func ExampleTokenURLSafe() {
	tok, _ := TokenURLSafe(24)
	fmt.Println(len(tok))
	// Output: 32
}
