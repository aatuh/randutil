package nanoid

import "fmt"

func ExampleID() {
	id, err := ID()
	if err != nil {
		fmt.Println("error")
		return
	}

	parsed, err := Parse(id)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(len(id))
	fmt.Println(parsed == id)
	// Output:
	// 21
	// true
}
