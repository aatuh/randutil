package adapters

import (
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

// LockedSource returns a Source wrapper that serializes access to src.
// If src is nil, it returns nil.
func LockedSource(src core.Source) core.Source {
	if src == nil {
		return nil
	}
	return &lockedSource{src: src}
}
