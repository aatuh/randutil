package dist

import (
	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func ExampleGenerator() {
	gen := New(core.New(adapters.DeterministicSource([]byte("seed"))))
	_, _ = gen.Normal(0, 1)
}
