package dist

import (
	"errors"
	"io"
	"math"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/internal/testutil"
)

func newGen(chunks ...[]byte) *Generator {
	return New(core.New(testutil.NewSeqReader(chunks...)))
}

func TestBernoulli(t *testing.T) {
	gen := New(nil)
	if _, err := gen.Bernoulli(-0.1); err == nil {
		t.Fatalf("expected error for p < 0")
	}
	if _, err := gen.Bernoulli(1.1); err == nil {
		t.Fatalf("expected error for p > 1")
	}
	for _, p := range []float64{math.NaN(), math.Inf(1), math.Inf(-1)} {
		if _, err := gen.Bernoulli(p); !errors.Is(err, core.ErrInvalidProbability) {
			t.Fatalf("Bernoulli(%v) error = %v want %v", p, err, core.ErrInvalidProbability)
		}
	}
	gen = newGen(testutil.Float64Bytes(0.2))
	v, err := gen.Bernoulli(0.5)
	if err != nil {
		t.Fatalf("Bernoulli error: %v", err)
	}
	if !v {
		t.Fatalf("Bernoulli expected true for u=0.2 < 0.5")
	}
	gen = newGen(testutil.Float64Bytes(0.8))
	v, err = gen.Bernoulli(0.5)
	if err != nil {
		t.Fatalf("Bernoulli error: %v", err)
	}
	if v {
		t.Fatalf("Bernoulli expected false for u=0.8 >= 0.5")
	}
}

func TestBernoulliErrorPropagation(t *testing.T) {
	gen := New(core.New(testutil.ErrReader{Err: io.ErrUnexpectedEOF}))
	if _, err := gen.Bernoulli(0.5); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestCategorical(t *testing.T) {
	weights := []float64{0.2, 0.3, 0.5}
	gen := newGen(testutil.Float64Bytes(0.1))
	idx, err := gen.Categorical(weights)
	if err != nil {
		t.Fatalf("Categorical error: %v", err)
	}
	if idx != 0 {
		t.Fatalf("Categorical idx=%d want 0", idx)
	}
	gen = newGen(testutil.Float64Bytes(0.6))
	idx, err = gen.Categorical(weights)
	if err != nil {
		t.Fatalf("Categorical error: %v", err)
	}
	if idx != 2 {
		t.Fatalf("Categorical idx=%d want 2", idx)
	}
	if _, err := gen.Categorical([]float64{}); err == nil {
		t.Fatalf("expected error for empty weights")
	}
	if _, err := gen.Categorical([]float64{math.NaN()}); err == nil {
		t.Fatalf("expected error for NaN weight")
	}
	if _, err := gen.Categorical([]float64{-1}); err == nil {
		t.Fatalf("expected error for negative weight")
	}
	if _, err := gen.Categorical([]float64{math.MaxFloat64, math.MaxFloat64}); !errors.Is(err, core.ErrInvalidWeights) {
		t.Fatalf("expected invalid weights for overflowing sum, got %v", err)
	}
}

func TestCategoricalErrorPropagation(t *testing.T) {
	weights := []float64{1}
	gen := New(core.New(testutil.ErrReader{Err: io.ErrUnexpectedEOF}))
	if _, err := gen.Categorical(weights); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestExponential(t *testing.T) {
	gen := New(nil)
	if _, err := gen.Exponential(0); err == nil {
		t.Fatalf("expected error for lambda <= 0")
	}
	gen = newGen(testutil.Float64Bytes(0.5))
	v, err := gen.Exponential(2)
	if err != nil {
		t.Fatalf("Exponential error: %v", err)
	}
	want := -math.Log(0.5) / 2
	if math.Abs(v-want) > 1e-12 {
		t.Fatalf("Exponential = %f want %f", v, want)
	}
}

func TestDistributionsRejectNonFiniteParameters(t *testing.T) {
	gen := New(nil)
	nonFinite := []float64{math.NaN(), math.Inf(1), math.Inf(-1)}

	for _, v := range nonFinite {
		if _, err := gen.Exponential(v); err == nil {
			t.Fatalf("Exponential(%v) error = nil", v)
		}
		if _, err := gen.Poisson(v); err == nil {
			t.Fatalf("Poisson(%v) error = nil", v)
		}
	}

	for _, tc := range []struct {
		name string
		min  float64
		max  float64
	}{
		{name: "nan min", min: math.NaN(), max: 1},
		{name: "nan max", min: 0, max: math.NaN()},
		{name: "negative infinity min", min: math.Inf(-1), max: 1},
		{name: "positive infinity max", min: 0, max: math.Inf(1)},
	} {
		if _, err := gen.Uniform(tc.min, tc.max); err == nil {
			t.Fatalf("Uniform %s error = nil", tc.name)
		}
	}

	for _, v := range nonFinite {
		if _, err := gen.Normal(v, 1); err == nil {
			t.Fatalf("Normal(%v, 1) error = nil", v)
		}
		if _, err := gen.Normal(0, v); err == nil {
			t.Fatalf("Normal(0, %v) error = nil", v)
		}
		if _, err := gen.Gamma(v, 1); err == nil {
			t.Fatalf("Gamma(%v, 1) error = nil", v)
		}
		if _, err := gen.Gamma(1, v); err == nil {
			t.Fatalf("Gamma(1, %v) error = nil", v)
		}
	}

	for _, tc := range []struct {
		name string
		s    float64
		v    float64
	}{
		{name: "nan s", s: math.NaN(), v: 1},
		{name: "infinite s", s: math.Inf(1), v: 1},
		{name: "nan v", s: 1, v: math.NaN()},
		{name: "infinite v", s: 1, v: math.Inf(1)},
	} {
		if _, err := gen.Zipf(tc.s, tc.v, 10); err == nil {
			t.Fatalf("Zipf %s error = nil", tc.name)
		}
	}
}

func TestExponentialErrorPropagation(t *testing.T) {
	gen := New(core.New(testutil.ErrReader{Err: io.ErrUnexpectedEOF}))
	if _, err := gen.Exponential(1); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}

func TestNormal(t *testing.T) {
	gen := New(nil)
	if _, err := gen.Normal(0, -1); err == nil {
		t.Fatalf("expected error for negative stddev")
	}
	gen = newGen(testutil.Float64Bytes(0.5), testutil.Float64Bytes(0.25))
	v, err := gen.Normal(2, 3)
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
	v, err = gen.Normal(5, 0)
	if err != nil {
		t.Fatalf("Normal zero stddev error: %v", err)
	}
	if v != 5 {
		t.Fatalf("Normal with zero stddev returned %f want 5", v)
	}
}

func TestNormalErrorPropagation(t *testing.T) {
	gen := New(core.New(testutil.ErrReader{Err: io.ErrUnexpectedEOF}))
	if _, err := gen.Normal(0, 1); err == nil {
		t.Fatalf("expected error from entropy source")
	}
}
