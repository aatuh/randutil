//go:build !randutil_policy
// +build !randutil_policy

package randutil

import (
	"fmt"
	"time"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/uuid"
)

func ExampleRand() {
	r := Default()
	b, err := r.Numeric.Bytes(8)
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(len(b))
	// Output: 8
}

func ExampleCollection() {
	src, err := adapters.DeterministicSourceWithLabel([]byte("seed"), "collection")
	if err != nil {
		fmt.Println("error")
		return
	}
	r := New(src)
	items := []string{"a", "b", "c", "d"}
	_ = Collection[string](r).Shuffle(items)
	fmt.Println(items)
	// Output: [b c a d]
}

func ExampleNew_deterministic() {
	seed := []byte("seed")
	fixed := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	src, err := adapters.DeterministicSourceWithLabel(seed, "uuid")
	if err != nil {
		fmt.Println("error")
		return
	}
	r := New(src)
	gen := uuid.NewWithClock(r.Core, func() time.Time { return fixed })
	u, err := gen.V7()
	if err != nil {
		fmt.Println("error")
		return
	}
	fmt.Println(u)
	// Output: 018cc820-d888-7a3b-ae5c-d5bc8d457263
}

func ExampleWorkspace() {
	root := DeterministicRoot([]byte("fixture seed"))
	ws := NewWorkspace(root)
	defer ws.Close()

	sessions, err := ws.Rand("sessions")
	if err != nil {
		fmt.Println("error")
		return
	}
	analytics, err := ws.Rand("analytics")
	if err != nil {
		fmt.Println("error")
		return
	}

	sessionID, err := sessions.String.String(8)
	if err != nil {
		fmt.Println("error")
		return
	}
	analyticsID, err := analytics.String.String(8)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(sessionID)
	fmt.Println(analyticsID)
	// Output:
	// qqu18myj
	// 6j35pnl8
}

func ExampleDerive() {
	r, err := Derive([]byte("fixture seed"), "reports")
	if err != nil {
		fmt.Println("error")
		return
	}

	code, err := r.String.String(10)
	if err != nil {
		fmt.Println("error")
		return
	}

	fmt.Println(code)
	// Output: jwupnx8zl7
}
