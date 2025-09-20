package dist

import (
	"io"
	"math"
	"testing"

	"github.com/aatuh/randutil/core"
	"github.com/aatuh/randutil/internal/testutil"
)

func withEntropy(t *testing.T, chunks ...[]byte) {
	t.Helper()
	core.SetSource(testutil.NewSeqReader(chunks...))
	t.Cleanup(core.ResetSource)
}

func TestBernoulli(t *testing.T) {
	if _, err := Bernoulli(-0.1); err == nil {
		t.Fatalf("expected error for p < 0")
	}
	if _, err := Bernoulli(1.1); err == nil {
		t.Fatalf("expected error for p > 1")
	}
	withEntropy(t, testutil.Float64Bytes(0.2))
	v, err := Bernoulli(0.5)
	if err != nil {
		t.Fatalf("Bernoulli error: %v", err)
	}
	if !v {
		t.Fatalf("Bernoulli expected true for u=0.2 < 0.5")
	}
	withEntropy(t, testutil.Float64Bytes(0.8))
	v, err = Bernoulli(0.5)
	if err != nil {
		t.Fatalf("Bernoulli error: %v", err)
	}
	if v {
		t.Fatalf("Bernoulli expected false for u=0.8 >= 0.5")
	}
}

func TestBernoulliErrorPropagation(t *testing.T) {
	core.SetSource(testutil.ErrReader{Err: io.ErrUnexpectedEOF})
	t.Cleanup(core.ResetSource)
	if _, err := Bernoulli(0.5); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestCategorical(t *testing.T) {
	weights := []float64{0.2, 0.3, 0.5}
	withEntropy(t, testutil.Float64Bytes(0.1))
	idx, err := Categorical(weights)
	if err != nil {
		t.Fatalf("Categorical error: %v", err)
	}
	if idx != 0 {
		t.Fatalf("Categorical idx=%d want 0", idx)
	}
	withEntropy(t, testutil.Float64Bytes(0.6))
	idx, err = Categorical(weights)
	if err != nil {
		t.Fatalf("Categorical error: %v", err)
	}
	if idx != 2 {
		t.Fatalf("Categorical idx=%d want 2", idx)
	}
	if _, err := Categorical([]float64{}); err == nil {
		t.Fatalf("expected error for empty weights")
	}
	if _, err := Categorical([]float64{math.NaN()}); err == nil {
		t.Fatalf("expected error for NaN weight")
	}
	if _, err := Categorical([]float64{-1}); err == nil {
		t.Fatalf("expected error for negative weight")
	}
}

func TestCategoricalErrorPropagation(t *testing.T) {
	weights := []float64{1}
	core.SetSource(testutil.ErrReader{Err: io.ErrUnexpectedEOF})
	t.Cleanup(core.ResetSource)
	if _, err := Categorical(weights); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestExponential(t *testing.T) {
	if _, err := Exponential(0); err == nil {
		t.Fatalf("expected error for lambda <= 0")
	}
	withEntropy(t, testutil.Float64Bytes(0.5))
	v, err := Exponential(2)
	if err != nil {
		t.Fatalf("Exponential error: %v", err)
	}
	want := -math.Log(0.5) / 2
	if math.Abs(v-want) > 1e-12 {
		t.Fatalf("Exponential = %f want %f", v, want)
	}
}

func TestExponentialErrorPropagation(t *testing.T) {
	core.SetSource(testutil.ErrReader{Err: io.ErrUnexpectedEOF})
	t.Cleanup(core.ResetSource)
	if _, err := Exponential(1); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestNormal(t *testing.T) {
	if _, err := Normal(0, -1); err == nil {
		t.Fatalf("expected error for negative stddev")
	}
	withEntropy(t, testutil.Float64Bytes(0.5), testutil.Float64Bytes(0.25))
	v, err := Normal(2, 3)
	if err != nil {
		t.Fatalf("Normal error: %v", err)
	}
	u1 := 0.5
	u2 := 0.25
	r := math.Sqrt(-2 * math.Log(u1))
	theta := 2 * math.Pi * u2
	want := 2 + 3*(r*math.Cos(theta))
	if math.Abs(v-want) > 1e-9 {
		t.Fatalf("Normal = %f want %f", v, want)
	}
	v, err = Normal(5, 0)
	if err != nil {
		t.Fatalf("Normal zero stddev error: %v", err)
	}
	if v != 5 {
		t.Fatalf("Normal with zero stddev returned %f want 5", v)
	}
}

func TestNormalErrorPropagation(t *testing.T) {
	core.SetSource(testutil.ErrReader{Err: io.ErrUnexpectedEOF})
	t.Cleanup(core.ResetSource)
	if _, err := Normal(0, 1); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}
