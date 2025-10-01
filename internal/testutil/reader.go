package testutil

import (
	"encoding/binary"
	"errors"
	"io"
	"math"
	"sync"
	"testing"

	"github.com/aatuh/randutil/v2/core"
)

// SeqReader feeds deterministic entropy to crypto consumers. It returns
// each configured byte in order and, once exhausted, keeps returning the
// final byte so callers never see EOF mid-read.
type SeqReader struct {
	mu   sync.Mutex
	data []byte
	pos  int
}

// NewSeqReader builds a SeqReader from one or more byte chunks.
func NewSeqReader(chunks ...[]byte) *SeqReader {
	total := 0
	for _, chunk := range chunks {
		total += len(chunk)
	}
	buf := make([]byte, 0, total)
	for _, chunk := range chunks {
		buf = append(buf, chunk...)
	}
	return &SeqReader{data: buf}
}

// Read implements io.Reader.
func (r *SeqReader) Read(p []byte) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if len(p) == 0 {
		return 0, nil
	}
	if len(r.data) == 0 {
		for i := range p {
			p[i] = 0
		}
		return len(p), nil
	}
	last := r.data[len(r.data)-1]
	for i := range p {
		if r.pos < len(r.data) {
			p[i] = r.data[r.pos]
			r.pos++
		} else {
			p[i] = last
		}
	}
	return len(p), nil
}

// ErrReader always fails with the provided error.
type ErrReader struct {
	Err error
}

// Read implements io.Reader.
func (e ErrReader) Read(p []byte) (int, error) {
	if e.Err == nil {
		return 0, errors.New("ErrReader: nil error supplied")
	}
	return 0, e.Err
}

const floatDenom = 1 << 53

// Float64Bytes returns an 8-byte little-endian payload that numeric.Float64
// will decode into the provided value. v is clamped to [0, 1).
func Float64Bytes(v float64) []byte {
	if math.IsNaN(v) {
		v = 0
	}
	if v < 0 {
		v = 0
	}
	if v >= 1 {
		v = math.Nextafter(1, 0)
	}
	u := uint64(v * floatDenom)
	raw := u << 11
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], raw)
	return buf[:]
}

// Uint64Bytes returns the little-endian encoding of u.
func Uint64Bytes(u uint64) []byte {
	var buf [8]byte
	binary.LittleEndian.PutUint64(buf[:], u)
	return buf[:]
}

// WithSource sets the entropy source for testing and restores it on cleanup.
func WithSource(t *testing.T, r io.Reader) {
	t.Helper()
	core.SetSource(r)
	t.Cleanup(core.ResetSource)
}
