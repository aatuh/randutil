package randtime

import (
	"fmt"
	"time"
)

func ExampleDatetime() {
	t, _ := Datetime()
	fmt.Println(t.Location() == time.UTC)
	// Output: true
}
