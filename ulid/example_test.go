package ulid

import (
	"fmt"
)

func ExampleGenerator_ULID() {
	gen := Default()

	id, err := gen.ULID()
	if err != nil {
		fmt.Println("error")
		return
	}
	parsed, err := Parse(id.String())
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(len(id))
	fmt.Println(parsed == id)
	// Output:
	// 26
	// true
}
