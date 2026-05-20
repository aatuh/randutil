//go:build !randutil_policy
// +build !randutil_policy

package randutil

import (
	"bytes"
	"errors"
	"sync"
	"testing"

	"github.com/aatuh/randutil/v2/core"
)

func TestWorkspaceStreamDeterministic(t *testing.T) {
	seed := []byte("seed")
	ws1 := NewWorkspace(DeterministicRoot(seed))
	ws2 := NewWorkspace(DeterministicRoot(seed))

	gen1, err := ws1.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	gen1b, err := ws1.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	if gen1 != gen1b {
		t.Fatalf("expected cached stream for label")
	}

	gen2, err := ws2.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	got, err := gen1.Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	want, err := gen2.Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if !bytes.Equal(got, want) {
		t.Fatalf("streams mismatch for same label")
	}

	gen3, err := ws1.Stream("beta")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	diff, err := gen3.Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if bytes.Equal(got, diff) {
		t.Fatalf("different labels produced identical output")
	}
}

func TestWorkspaceStreamConcurrency(t *testing.T) {
	ws := NewWorkspace(DeterministicRoot([]byte("seed")))
	const workers = 32
	var wg sync.WaitGroup
	gens := make(chan *core.Generator, workers)
	errs := make(chan error, workers)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			gen, err := ws.Stream("alpha")
			if err != nil {
				errs <- err
				return
			}
			gens <- gen
		}()
	}
	wg.Wait()
	close(errs)
	close(gens)

	for err := range errs {
		t.Fatalf("Stream error: %v", err)
	}
	var first *core.Generator
	for gen := range gens {
		if first == nil {
			first = gen
			continue
		}
		if gen != first {
			t.Fatalf("expected cached stream for label")
		}
	}
}

func TestWorkspaceUsageAndClose(t *testing.T) {
	ws := NewWorkspace(DeterministicRoot([]byte("seed")))
	gen, err := ws.Stream("alpha")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	if _, err := gen.Bytes(16); err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	used, ok := ws.Usage("alpha")
	if !ok {
		t.Fatalf("expected usage for label")
	}
	if used != 16 {
		t.Fatalf("usage=%d want 16", used)
	}
	if err := ws.Close(); err != nil {
		t.Fatalf("Close error: %v", err)
	}
	if _, err := ws.Stream("alpha"); !errors.Is(err, core.ErrWorkspaceClosed) {
		t.Fatalf("expected ErrWorkspaceClosed, got %v", err)
	}
}

func TestWorkspaceSubDeterministic(t *testing.T) {
	seed := []byte("seed")
	ws1 := NewWorkspace(DeterministicRoot(seed))
	ws2 := NewWorkspace(DeterministicRoot(seed))
	sub1, err := ws1.Sub("payments")
	if err != nil {
		t.Fatalf("Sub error: %v", err)
	}
	sub2, err := ws2.Sub("payments")
	if err != nil {
		t.Fatalf("Sub error: %v", err)
	}
	gen1, err := sub1.Stream("nonces")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	gen2, err := sub2.Stream("nonces")
	if err != nil {
		t.Fatalf("Stream error: %v", err)
	}
	b1, err := gen1.Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	b2, err := gen2.Bytes(32)
	if err != nil {
		t.Fatalf("Bytes error: %v", err)
	}
	if !bytes.Equal(b1, b2) {
		t.Fatalf("sub-workspaces mismatch")
	}
}
