package uuid

import (
	"errors"
	"strings"
	"testing"
	stdtime "time"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestV4DeterministicEntropy(t *testing.T) {
	data := make([]byte, 16)
	for i := range data {
		data[i] = byte(i)
	}
	gen := New(core.New(testutil.NewSeqReader(data)))
	u, err := gen.V4()
	if err != nil {
		t.Fatalf("V4 error: %v", err)
	}
	if !strings.HasPrefix(string(u), "00010203") {
		t.Fatalf("unexpected prefix: %s", u)
	}
	bytes, err := u.Bytes()
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if (bytes[6] >> 4) != 4 {
		t.Fatalf("version nibble = %x want 4", bytes[6]>>4)
	}
	if bytes[8]&0xc0 != 0x80 {
		t.Fatalf("variant bits incorrect: %x", bytes[8])
	}
}

func TestV4ErrorPropagation(t *testing.T) {
	gen := New(core.New(testutil.ErrReader{Err: errors.New("entropy failure")}))
	if _, err := gen.V4(); err == nil {
		t.Fatalf("expected error when entropy fails")
	}
	defer func() {
		if recover() == nil {
			t.Fatalf("MustV4 did not panic on entropy failure")
		}
	}()
	gen.MustV4()
}

func TestV7Structure(t *testing.T) {
	gen := New(core.New(testutil.NewSeqReader(make([]byte, 16))))
	before := stdtime.Now().UTC()
	u, err := gen.V7()
	if err != nil {
		t.Fatalf("V7 error: %v", err)
	}
	bytes, err := u.Bytes()
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if (bytes[6] >> 4) != 7 {
		t.Fatalf("version nibble = %x want 7", bytes[6]>>4)
	}
	if bytes[8]&0xc0 != 0x80 {
		t.Fatalf("variant bits incorrect: %x", bytes[8])
	}
	ts := int64(bytes[0])<<40 | int64(bytes[1])<<32 | int64(bytes[2])<<24 |
		int64(bytes[3])<<16 | int64(bytes[4])<<8 | int64(bytes[5])
	now := stdtime.Now().UTC()
	if ts < before.UnixMilli() || ts > now.UnixMilli() {
		t.Fatalf("timestamp %d not within bounds [%d,%d]", ts, before.UnixMilli(), now.UnixMilli())
	}
}

func TestV7WithClock(t *testing.T) {
	fixed := stdtime.Date(2024, 1, 2, 3, 4, 5, 0, stdtime.UTC)
	gen := NewWithClock(core.New(testutil.NewSeqReader(make([]byte, 16))),
		func() stdtime.Time { return fixed })
	u, err := gen.V7()
	if err != nil {
		t.Fatalf("V7 error: %v", err)
	}
	bytes, err := u.Bytes()
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	ts := int64(bytes[0])<<40 | int64(bytes[1])<<32 | int64(bytes[2])<<24 |
		int64(bytes[3])<<16 | int64(bytes[4])<<8 | int64(bytes[5])
	if ts != fixed.UnixMilli() {
		t.Fatalf("timestamp %d want %d", ts, fixed.UnixMilli())
	}
}

func TestParseAndMustParse(t *testing.T) {
	u, err := Parse("A8098C1A-F86E-11DA-BDBF-10B96E4EF00D")
	if err != nil {
		t.Fatalf("Parse error: %v", err)
	}
	if string(u) != strings.ToLower("A8098C1A-F86E-11DA-BDBF-10B96E4EF00D") {
		t.Fatalf("Parse did not lowercase result: %s", u)
	}
	if MustParse("00000000-0000-0000-0000-000000000000").IsNil() != true {
		t.Fatalf("MustParse nil UUID not detected")
	}
	defer func() {
		if recover() == nil {
			t.Fatalf("MustParse did not panic on invalid input")
		}
	}()
	MustParse("not-a-uuid")
}

func TestUUIDBytesValidation(t *testing.T) {
	u := UUID("00000000-0000-0000-0000-000000000000")
	_, err := u.Bytes()
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if _, err := UUID("00000000-0000-0000-0000-00000000000G").Bytes(); err == nil {
		t.Fatalf("expected error for invalid hex")
	}
	if _, err := UUID("00000000-0000-0000-0000-000000000000").Bytes(); err != nil {
		t.Fatalf("valid UUID unexpectedly failed: %v", err)
	}
}

func TestNilHelpers(t *testing.T) {
	if !Nil().IsNil() {
		t.Fatalf("Nil().IsNil() should be true")
	}
	if Nil().String() != "00000000-0000-0000-0000-000000000000" {
		t.Fatalf("Nil string mismatch: %s", Nil())
	}
}
