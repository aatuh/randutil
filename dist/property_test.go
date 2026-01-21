package dist

import (
	"math"
	"testing"

	"github.com/aatuh/randutil/v2/adapters"
	"github.com/aatuh/randutil/v2/core"
)

func TestDistributionStats(t *testing.T) {
	const n = 20000
	cases := []struct {
		name         string
		expectedMean float64
		expectedVar  float64
		sample       func(*Generator) (float64, error)
	}{
		{
			name:         "normal",
			expectedMean: 1.5,
			expectedVar:  4.0,
			sample: func(g *Generator) (float64, error) {
				return g.Normal(1.5, 2.0)
			},
		},
		{
			name:         "exponential",
			expectedMean: 0.5,
			expectedVar:  0.25,
			sample: func(g *Generator) (float64, error) {
				return g.Exponential(2.0)
			},
		},
		{
			name:         "gamma",
			expectedMean: 2.5 / 1.3,
			expectedVar:  2.5 / (1.3 * 1.3),
			sample: func(g *Generator) (float64, error) {
				return g.Gamma(2.5, 1.3)
			},
		},
		{
			name:         "uniform",
			expectedMean: 1.0,
			expectedVar:  64.0 / 12.0,
			sample: func(g *Generator) (float64, error) {
				return g.Uniform(-3, 5)
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			gen := New(core.New(adapters.DeterministicSource([]byte(tc.name))))
			mean, variance := sampleStats(t, n, func() (float64, error) {
				return tc.sample(gen)
			})
			assertStats(t, mean, variance, tc.expectedMean, tc.expectedVar, n)
		})
	}

	t.Run("poisson", func(t *testing.T) {
		gen := New(core.New(adapters.DeterministicSource([]byte("poisson"))))
		mean, variance := sampleStats(t, n, func() (float64, error) {
			v, err := gen.Poisson(12)
			return float64(v), err
		})
		assertStats(t, mean, variance, 12, 12, n)
	})
}

func TestDeterministicSequence(t *testing.T) {
	seed := []byte("deterministic")
	g1 := New(core.New(adapters.DeterministicSource(seed)))
	g2 := New(core.New(adapters.DeterministicSource(seed)))
	for i := 0; i < 10; i++ {
		v1, err := g1.Normal(0, 1)
		if err != nil {
			t.Fatalf("Normal error: %v", err)
		}
		v2, err := g2.Normal(0, 1)
		if err != nil {
			t.Fatalf("Normal error: %v", err)
		}
		if v1 != v2 {
			t.Fatalf("deterministic mismatch at %d: %f vs %f", i, v1, v2)
		}
	}
}

func sampleStats(t *testing.T, n int, fn func() (float64, error)) (float64, float64) {
	mean := 0.0
	m2 := 0.0
	for i := 1; i <= n; i++ {
		x, err := fn()
		if err != nil {
			t.Fatalf("sample error: %v", err)
		}
		delta := x - mean
		mean += delta / float64(i)
		m2 += delta * (x - mean)
	}
	variance := m2 / float64(n-1)
	return mean, variance
}

func assertStats(t *testing.T, mean, variance, expectedMean, expectedVar float64, n int) {
	meanTol := 6 * math.Sqrt(expectedVar/float64(n))
	if meanTol < 1e-3 {
		meanTol = 1e-3
	}
	if math.Abs(mean-expectedMean) > meanTol {
		t.Fatalf("mean = %f want %f ± %f", mean, expectedMean, meanTol)
	}
	varTol := expectedVar * 0.15
	if varTol < 1e-3 {
		varTol = 1e-3
	}
	if math.Abs(variance-expectedVar) > varTol {
		t.Fatalf("variance = %f want %f ± %f", variance, expectedVar, varTol)
	}
}
