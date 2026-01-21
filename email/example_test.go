package email

import (
	"fmt"
	"strings"
)

func ExampleEmail() {
	s, _ := Email(Options{TLD: "org"})
	fmt.Println(strings.HasSuffix(s, ".org"))
	// Output: true
}
