package set_test

import (
	"testing"

	"github.com/deckarep/golang-set"
	"github.com/nlandolfi/set"
)

// Benchmarks the set implementation in package set
func BenchmarkSet(b *testing.B) {
	s := set.New()

	for n := 0; n < b.N; n++ {
		ok := s.Contains(n)
		if !ok {
			s.Add(n)
		}
	}
}

// Benchmarks managing membership with the primitive
// golang map
func BenchmarkSetPrimitive(b *testing.B) {
	s := make(map[int]bool)

	for n := 0; n < b.N; n++ {
		_, ok := s[n]
		if !ok {
			s[n] = true
		}
	}
}

// Benchmarks the thread unsafe version of the mapset,
// the most prominent alternative
func BenchmarkGolangSet(b *testing.B) {
	s := mapset.NewThreadUnsafeSet()

	for n := 0; n < b.N; n++ {
		ok := s.Contains(n)
		if !ok {
			s.Add(n)
		}
	}
}
