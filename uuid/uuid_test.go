package uuid

import (
	"errors"
	"strings"
	"testing"
	stdtime "time"

	"github.com/aatuh/randutil/core"
	"github.com/aatuh/randutil/internal/testutil"
)

func TestV4DeterministicEntropy(t *testing.T) {
	data := make([]byte, 16)
	for i := range data {
		data[i] = byte(i)
	}
	core.SetSource(testutil.NewSeqReader(data))
	t.Cleanup(core.ResetSource)
	u, err := V4()
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
	core.SetSource(testutil.ErrReader{Err: errors.New("entropy failure")})
	t.Cleanup(core.ResetSource)
	if _, err := V4(); err == nil {
		t.Fatalf("expected error when entropy fails")
	}
	defer func() {
		if recover() == nil {
			t.Fatalf("MustV4 did not panic on entropy failure")
		}
	}()
	MustV4()
}

func TestV7Structure(t *testing.T) {
	core.SetSource(testutil.NewSeqReader(make([]byte, 16)))
	t.Cleanup(core.ResetSource)
	before := stdtime.Now().UTC()
	u, err := V7()
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
	b, err := u.Bytes()
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if len(b) != 16 {
		t.Fatalf("Bytes length = %d want 16", len(b))
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
