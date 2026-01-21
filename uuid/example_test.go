package uuid

import "fmt"

func ExampleV4() {
	u, _ := V4()
	fmt.Println(len(u) == 36)
	// Output: true
}
