package adapters

import (
	"io"
	"sync"

	"github.com/aatuh/randutil/v2/core"
)

const defaultBufferSize = 4096

type bufferedSource struct {
	mu  sync.Mutex
	src core.Source
	buf []byte
	pos int
	n   int
}

// BufferedSource wraps src with a buffer to amortize small reads.
// If src is nil, it returns nil.
func BufferedSource(src core.Source) core.Source {
	return BufferedSourceWithSize(src, defaultBufferSize)
}

// BufferedSourceWithSize wraps src with an internal buffer of size bytes.
// If size <= 0, a default buffer size is used. If src is nil, it returns nil.
func BufferedSourceWithSize(src core.Source, size int) core.Source {
	if src == nil {
		return nil
	}
	if size <= 0 {
		size = defaultBufferSize
	}
	return &bufferedSource{
		src: src,
		buf: make([]byte, size),
	}
}

func (b *bufferedSource) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	if b == nil {
		return 0, core.ErrSourceClosed
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	if b.src == nil {
		return 0, core.ErrSourceClosed
	}

	total := 0
	for total < len(p) {
		if b.pos >= b.n {
			n, err := io.ReadFull(b.src, b.buf)
			if err != nil {
				if total > 0 {
					return total, err
				}
				return 0, err
			}
			b.pos = 0
			b.n = n
		}
		copied := copy(p[total:], b.buf[b.pos:b.n])
		b.pos += copied
		total += copied
	}
	return total, nil
}

func (b *bufferedSource) Close() error {
	if b == nil {
		return nil
	}
	b.mu.Lock()
	defer b.mu.Unlock()
	core.Zero(b.buf)
	b.pos = 0
	b.n = 0
	src := b.src
	b.src = nil
	if closer, ok := src.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}
