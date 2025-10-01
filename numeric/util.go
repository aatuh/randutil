package numeric

import (
	"github.com/aatuh/randutil/v2/core"
)

// def is a package-level zero-value generator that uses the live proxy source.
var def core.Generator // zero-value uses core.Reader()
