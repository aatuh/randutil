package core

import (
	"bytes"
	"crypto/rand"
	"io"
	"testing"
)

type countingReader struct {
	buf []byte
}

func (c *countingReader) Read(p []byte) (int, error) {
	n := copy(p, c.buf)
	if n < len(p) {
		for i := n; i < len(p); i++ {
			p[i] = byte(i)
		}
	}
	return len(p), nil
}

func TestSetAndGetSource(t *testing.T) {
	defaultSrc := GetSource()
	if defaultSrc == nil {
		t.Fatal("default source must not be nil")
	}
	fake := &countingReader{buf: []byte{1, 2, 3, 4}}
	SetSource(fake)
	t.Cleanup(ResetSource)
	if got := GetSource(); got != fake {
		t.Fatalf("GetSource returned %T, want %T", got, fake)
	}
}

func TestSetSourceNilDoesNotOverwrite(t *testing.T) {
	orig := GetSource()
	SetSource(nil)
	if got := GetSource(); got != orig {
		t.Fatalf("SetSource(nil) changed source: got %T want %T", got, orig)
	}
}

func TestResetSourceRestoresCryptoRand(t *testing.T) {
	fake := bytes.NewBuffer([]byte{1, 2, 3})
	SetSource(fake)
	ResetSource()
	if GetSource() != rand.Reader {
		t.Fatalf("ResetSource did not restore crypto/rand Reader")
	}
}

func TestReaderDelegatesToCurrentSource(t *testing.T) {
	fake := &countingReader{buf: []byte{0xaa}}
	SetSource(fake)
	t.Cleanup(ResetSource)
	buf := make([]byte, 2)
	n, err := Reader().Read(buf)
	if err != nil {
		t.Fatalf("Reader returned error: %v", err)
	}
	if n != len(buf) {
		t.Fatalf("Reader returned %d bytes, want %d", n, len(buf))
	}
	expected := []byte{0xaa, 0x01}
	if !bytes.Equal(buf, expected) {
		t.Fatalf("Reader bytes = %v want %v", buf, expected)
	}
}

func TestReaderConcurrentAccess(t *testing.T) {
	fake := &countingReader{buf: []byte{0x42}}
	SetSource(fake)
	t.Cleanup(ResetSource)
	done := make(chan struct{})
	go func() {
		buf := make([]byte, 8)
		_, _ = Reader().Read(buf)
		close(done)
	}()
	if _, err := io.ReadAll(io.LimitReader(Reader(), 1)); err != nil {
		t.Fatalf("Reader streaming failed: %v", err)
	}
	<-done
}
