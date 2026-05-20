package adapters

import (
	"io"
	"sync"

	"github.com/aatuh/randutil/v2/core"
)

type lockedSource struct {
	mu  sync.Mutex
	src core.Source
}

func (l *lockedSource) Read(p []byte) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.src.Read(p)
}

func (l *lockedSource) Close() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	if closer, ok := l.src.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

// LockedSource returns a Source wrapper that serializes access to src.
// If src is nil, it returns nil.
func LockedSource(src core.Source) core.Source {
	if src == nil {
		return nil
	}
	return &lockedSource{src: src}
}
