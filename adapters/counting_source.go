package adapters

import (
	"io"
	"sync/atomic"

	"github.com/aatuh/randutil/v2/core"
)

// CountingSource wraps an entropy source and tracks the total bytes read.
type CountingSource struct {
	src   core.Source
	count atomic.Uint64
	hook  func(delta uint64)
}

// NewCountingSource returns a CountingSource that wraps src.
// If src is nil, it returns nil.
func NewCountingSource(src core.Source, hook func(delta uint64)) *CountingSource {
	if src == nil {
		return nil
	}
	return &CountingSource{
		src:  src,
		hook: hook,
	}
}

// Read reads from the underlying source and updates the byte count.
func (c *CountingSource) Read(p []byte) (int, error) {
	if c == nil || c.src == nil {
		return 0, core.ErrSourceClosed
	}
	n, err := c.src.Read(p)
	if n > 0 {
		// #nosec G115 -- n is a positive byte count returned by Read; int fits in uint64 on supported Go platforms.
		delta := uint64(n)
		c.count.Add(delta)
		if c.hook != nil {
			c.hook(delta)
		}
	}
	return n, err
}

// Count returns the total number of bytes read from the source.
func (c *CountingSource) Count() uint64 {
	if c == nil {
		return 0
	}
	return c.count.Load()
}

// Close closes the underlying source if it is closable.
func (c *CountingSource) Close() error {
	if c == nil {
		return nil
	}
	if closer, ok := c.src.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
