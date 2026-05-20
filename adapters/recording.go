package adapters

import (
	"io"
	"sync"

	"github.com/aatuh/randutil/v2/core"
)

// Recorder wraps a source and records all bytes read from it.
type Recorder struct {
	mu     sync.Mutex
	src    core.Source
	buf    []byte
	closed bool
}

// NewRecorder returns a Recorder that wraps src. If src is nil, it returns nil.
func NewRecorder(src core.Source) *Recorder {
	if src == nil {
		return nil
	}
	return &Recorder{src: src}
}

// Read reads from the underlying source and appends the bytes to the record.
func (r *Recorder) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed || r.src == nil {
		return 0, core.ErrSourceClosed
	}
	n, err := r.src.Read(p)
	if n > 0 {
		r.buf = append(r.buf, p[:n]...)
	}
	return n, err
}

// Bytes returns a copy of the recorded bytes.
func (r *Recorder) Bytes() []byte {
	r.mu.Lock()
	defer r.mu.Unlock()
	out := make([]byte, len(r.buf))
	copy(out, r.buf)
	return out
}

// Reset clears the recorded bytes.
func (r *Recorder) Reset() {
	r.mu.Lock()
	defer r.mu.Unlock()
	core.Zero(r.buf)
	r.buf = r.buf[:0]
}

// Replay returns a replay source for the recorded bytes.
func (r *Recorder) Replay() core.Source {
	return ReplaySource(r.Bytes())
}

// Close closes the underlying source if it is closable and zeroes the buffer.
func (r *Recorder) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	r.closed = true
	core.Zero(r.buf)
	r.buf = nil
	if closer, ok := r.src.(io.Closer); ok {
		return closer.Close()
	}
	return nil
}

type replaySource struct {
	mu     sync.Mutex
	data   []byte
	pos    int
	closed bool
}

// ReplaySource returns a source that replays the provided bytes in order.
func ReplaySource(data []byte) core.Source {
	if data == nil {
		data = []byte{}
	}
	copied := make([]byte, len(data))
	copy(copied, data)
	return &replaySource{data: copied}
}

func (r *replaySource) Read(p []byte) (int, error) {
	if len(p) == 0 {
		return 0, nil
	}
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return 0, core.ErrSourceClosed
	}
	remaining := len(r.data) - r.pos
	if remaining <= 0 {
		return 0, io.EOF
	}
	if len(p) > remaining {
		copy(p, r.data[r.pos:])
		for i := remaining; i < len(p); i++ {
			p[i] = 0
		}
		r.pos = len(r.data)
		return remaining, io.ErrUnexpectedEOF
	}
	copy(p, r.data[r.pos:r.pos+len(p)])
	r.pos += len(p)
	return len(p), nil
}

func (r *replaySource) Close() error {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.closed {
		return nil
	}
	r.closed = true
	core.Zero(r.data)
	r.data = nil
	r.pos = 0
	return nil
}
