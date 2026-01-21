package randutil

import (
	"fmt"
	"time"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/uuid"
)

func ExampleRand() {
	r := Default()
	b := r.Numeric.MustBytes(8)
	fmt.Println(len(b))
	// Output: 8
}

func ExampleCollection() {
	r := New(adapters.DeterministicSourceWithLabel([]byte("seed"), "collection"))
	items := []string{"a", "b", "c", "d"}
	_ = Collection[string](r).Shuffle(items)
	fmt.Println(items)
	// Output: [b c a d]
}

func ExampleNew_deterministic() {
	seed := []byte("seed")
	fixed := time.Date(2024, 1, 2, 3, 4, 5, 0, time.UTC)
	r := New(adapters.DeterministicSourceWithLabel(seed, "uuid"))
	gen := uuid.NewWithClock(r.Core, func() time.Time { return fixed })
	fmt.Println(gen.MustV7())
	// Output: 018cc820-d888-7a3b-ae5c-d5bc8d457263
}
