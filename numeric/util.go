package numeric

import (
	"io"

	"github.com/aatuh/randutil/core"
)

// readFull wraps io.ReadFull with the active entropy source.
func readFull(p []byte) error {
	_, err := io.ReadFull(core.GetSource(), p)
	return err
}
