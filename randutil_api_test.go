package randutil

import (
	"bytes"
	crand "crypto/rand"
	"errors"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func TestNewWiresGeneratorsToSource(t *testing.T) {
	src := testutil.NewSeqReader([]byte{1, 2, 3, 4, 5, 6})
	r := New(src)
	assertRandReady(t, r)
	if r.Source() != src {
		t.Fatalf("Source() = %T, want original source", r.Source())
	}

	b, err := r.Numeric.Bytes(4)
	if err != nil {
		t.Fatalf("Numeric.Bytes error: %v", err)
	}
	if !bytes.Equal(b, []byte{1, 2, 3, 4}) {
		t.Fatalf("Numeric.Bytes = %v want [1 2 3 4]", b)
	}

	buf := make([]byte, 2)
	if err := r.Core.Fill(buf); err != nil {
		t.Fatalf("Core.Fill error: %v", err)
	}
	if !bytes.Equal(buf, []byte{5, 6}) {
		t.Fatalf("Core.Fill = %v want [5 6]", buf)
	}
}

func TestDefaultAndSecureUseCryptoRand(t *testing.T) {
	for name, r := range map[string]Rand{
		"Default": Default(),
		"Secure":  Secure(),
	} {
		assertRandReady(t, r)
		if r.Source() != crand.Reader {
			t.Fatalf("%s Source() = %T, want crypto/rand.Reader", name, r.Source())
		}
	}
}

func TestCollectionUsesRandCore(t *testing.T) {
	r := New(testutil.NewSeqReader(testutil.Uint64Bytes(1)))
	got, err := Collection[string](r).PickOne([]string{"a", "b", "c"})
	if err != nil {
		t.Fatalf("PickOne error: %v", err)
	}
	if got != "b" {
		t.Fatalf("PickOne = %q want b", got)
	}

	if _, err := Collection[int](Rand{}).PickOne([]int{1}); err != nil {
		t.Fatalf("Collection fallback PickOne error: %v", err)
	}
}

func TestDeriveAndDeriveRNGShareDerivation(t *testing.T) {
	seed := []byte("root-seed")
	label := "api-test"

	r, err := Derive(seed, label)
	if err != nil {
		t.Fatalf("Derive error: %v", err)
	}
	rng, err := DeriveRNG(seed, label)
	if err != nil {
		t.Fatalf("DeriveRNG error: %v", err)
	}

	fromRand, err := r.Core.Bytes(16)
	if err != nil {
		t.Fatalf("Rand bytes error: %v", err)
	}
	fromRNG, err := rng.Bytes(16)
	if err != nil {
		t.Fatalf("RNG bytes error: %v", err)
	}
	if !bytes.Equal(fromRand, fromRNG) {
		t.Fatalf("Derive and DeriveRNG produced different streams: %x vs %x", fromRand, fromRNG)
	}
}

func TestDeriveSeparatesLabels(t *testing.T) {
	seed := []byte("root-seed")
	first, err := Derive(seed, "alpha")
	if err != nil {
		t.Fatalf("Derive alpha error: %v", err)
	}
	repeated, err := Derive(seed, "alpha")
	if err != nil {
		t.Fatalf("Derive alpha repeat error: %v", err)
	}
	other, err := Derive(seed, "beta")
	if err != nil {
		t.Fatalf("Derive beta error: %v", err)
	}

	firstBytes, err := first.Core.Bytes(32)
	if err != nil {
		t.Fatalf("first Bytes error: %v", err)
	}
	repeatedBytes, err := repeated.Core.Bytes(32)
	if err != nil {
		t.Fatalf("repeated Bytes error: %v", err)
	}
	otherBytes, err := other.Core.Bytes(32)
	if err != nil {
		t.Fatalf("other Bytes error: %v", err)
	}
	if !bytes.Equal(firstBytes, repeatedBytes) {
		t.Fatalf("same seed and label produced different streams: %x vs %x", firstBytes, repeatedBytes)
	}
	if bytes.Equal(firstBytes, otherBytes) {
		t.Fatalf("different labels produced identical derived streams: %x", firstBytes)
	}
}

func TestSecureRootWithSourceUsesInjectedSeed(t *testing.T) {
	seed := []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23,
		24, 25, 26, 27, 28, 29, 30, 31,
	}
	root := SecureRootWithSource(testutil.NewSeqReader(seed))
	derived, err := root.Derive("alpha")
	if err != nil {
		t.Fatalf("root Derive error: %v", err)
	}
	wantSource, err := adapters.DeriveSource(seed, "alpha")
	if err != nil {
		t.Fatalf("DeriveSource error: %v", err)
	}

	got := make([]byte, 32)
	if _, err := derived.Read(got); err != nil {
		t.Fatalf("derived Read error: %v", err)
	}
	want := make([]byte, 32)
	if _, err := wantSource.Read(want); err != nil {
		t.Fatalf("want Read error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("SecureRootWithSource derived %x, want %x", got, want)
	}
}

func TestSecureRootCloseClosesSourceAndRejectsDerive(t *testing.T) {
	src := &closeTrackingSource{Source: testutil.NewSeqReader(make([]byte, 32))}
	root := SecureRootWithSource(src)
	if _, err := root.Derive("alpha"); err != nil {
		t.Fatalf("Derive before Close error: %v", err)
	}
	if err := root.(interface{ Close() error }).Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
	if !src.closed {
		t.Fatalf("SecureRoot Close did not close underlying source")
	}
	if _, err := root.Derive("alpha"); !errors.Is(err, core.ErrSourceClosed) {
		t.Fatalf("Derive after Close error = %v, want ErrSourceClosed", err)
	}
	if err := root.(interface{ Close() error }).Close(); err != nil {
		t.Fatalf("second Close error: %v", err)
	}
}

func TestFastWithSourceUsesProvidedSeedSource(t *testing.T) {
	seed := []byte{
		0, 1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14, 15,
		16, 17, 18, 19, 20, 21, 22, 23,
		24, 25, 26, 27, 28, 29, 30, 31,
	}
	r, err := FastWithSource(testutil.NewSeqReader(seed))
	if err != nil {
		t.Fatalf("FastWithSource error: %v", err)
	}
	direct, err := adapters.FastSourceWithSource(testutil.NewSeqReader(seed))
	if err != nil {
		t.Fatalf("FastSourceWithSource error: %v", err)
	}

	fromRand, err := r.Core.Bytes(16)
	if err != nil {
		t.Fatalf("Rand bytes error: %v", err)
	}
	fromSource := make([]byte, 16)
	if _, err := direct.Read(fromSource); err != nil {
		t.Fatalf("direct Read error: %v", err)
	}
	if !bytes.Equal(fromRand, fromSource) {
		t.Fatalf("FastWithSource output mismatch: %x vs %x", fromRand, fromSource)
	}
}

type closeTrackingSource struct {
	core.Source
	closed bool
}

func (s *closeTrackingSource) Close() error {
	s.closed = true
	return nil
}

func assertRandReady(t *testing.T, r Rand) {
	t.Helper()
	if r.Core == nil ||
		r.Numeric == nil ||
		r.Dist == nil ||
		r.String == nil ||
		r.UUID == nil ||
		r.Time == nil ||
		r.Email == nil ||
		r.NanoID == nil ||
		r.ULID == nil {
		t.Fatalf("Rand has nil generator: %#v", r)
	}
}
