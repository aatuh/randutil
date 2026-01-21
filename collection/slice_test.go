package collection

import (
	"errors"
	"sort"
	"testing"

	"github.com/aatuh/randutil/v2/core"
	"github.com/aatuh/randutil/v2/numeric"
)

func TestShufflePermutation(t *testing.T) {
	n := 100
	p := make([]int, n)
	for i := range p {
		p[i] = i
	}
	if err := Shuffle(p); err != nil {
		t.Fatalf("Shuffle error: %v", err)
	}
	seen := map[int]int{}
	for _, v := range p {
		seen[v]++
	}
	if len(seen) != n {
		t.Fatalf("not a permutation; unique=%d", len(seen))
	}
	for v, c := range seen {
		if c != 1 {
			t.Fatalf("value %d occurs %d times", v, c)
		}
	}
}

func TestSampleBasics(t *testing.T) {
	src := []int{1, 2, 3, 4, 5, 6}
	original := append([]int(nil), src...) // Keep a copy to check original wasn't modified
	k := 3
	got, err := Sample(src, k)
	if err != nil {
		t.Fatalf("Sample error: %v", err)
	}
	if len(got) != k {
		t.Fatalf("len(got)=%d want %d", len(got), k)
	}
	// Check items are from src and unique.
	cpy := append([]int(nil), src...)
	sort.Ints(cpy)
	seen := map[int]bool{}
	for _, v := range got {
		if seen[v] {
			t.Fatalf("duplicate in sample: %d", v)
		}
		seen[v] = true
		// binary search in cpy
		i := sort.SearchInts(cpy, v)
		if i >= len(cpy) || cpy[i] != v {
			t.Fatalf("value not from source: %d", v)
		}
	}
	// Check that original slice wasn't modified
	for i := range original {
		if src[i] != original[i] {
			t.Fatalf("Sample modified original slice at %d", i)
		}
	}
}

func TestPerm(t *testing.T) {
	src := make([]int, 50)
	for i := range src {
		src[i] = i
	}
	original := append([]int(nil), src...)
	p, err := Perm(src)
	if err != nil {
		t.Fatalf("Perm error: %v", err)
	}
	if len(p) != len(src) {
		t.Fatalf("perm length = %d want %d", len(p), len(src))
	}
	seen := map[int]bool{}
	for _, v := range p {
		if seen[v] {
			t.Fatalf("duplicate value in perm: %d", v)
		}
		seen[v] = true
	}
	for i, v := range src {
		if v != original[i] {
			t.Fatalf("Perm modified input slice at %d", i)
		}
	}
}

func TestUint64nInvalid(t *testing.T) {
	if _, err := numeric.Uint64n(0); err == nil {
		t.Fatal("expected error for n == 0")
	}
}

func TestPickOne(t *testing.T) {
	s := []int{1, 2, 3}
	seen := make(map[int]bool)
	for i := 0; i < 25; i++ {
		v, err := PickOne(s)
		if err != nil {
			t.Fatalf("PickOne error: %v", err)
		}
		seen[v] = true
	}
	if len(seen) == 0 {
		t.Fatalf("PickOne never returned any value")
	}
}

func TestPickOneEmpty(t *testing.T) {
	if _, err := PickOne([]int{}); err == nil {
		t.Fatal("expected error for empty slice")
	}
}

func TestSampleErrors(t *testing.T) {
	s := []int{1, 2, 3}
	if _, err := Sample(s, -1); !errors.Is(err, core.ErrNegativeLength) {
		t.Fatalf("Sample negative error = %v", err)
	}
	if _, err := Sample(s, 4); !errors.Is(err, core.ErrSampleTooLarge) {
		t.Fatalf("Sample oversize error = %v", err)
	}
}

func TestChoiceAndMust(t *testing.T) {
	v, err := Choice(10, 20, 30)
	if err != nil {
		t.Fatalf("Choice error: %v", err)
	}
	if !contains([]int{10, 20, 30}, v) {
		t.Fatalf("Choice returned unexpected value %d", v)
	}
	if MustChoice("a", "b") == "" {
		t.Fatal("MustChoice returned empty result")
	}
}

func TestMustPickPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("MustPickOne did not panic for empty slice")
		}
	}()
	MustPickOne([]int{})
}

func contains[T comparable](s []T, v T) bool {
	for _, x := range s {
		if x == v {
			return true
		}
	}
	return false
}
