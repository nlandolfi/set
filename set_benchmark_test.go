package set

import "testing"

func BenchmarkSet(b *testing.B) {
	s := New()

	for n := 0; n < b.N; n++ {
		ok := s.Contains(n)
		if !ok {
			s.Add(n)
		}
	}
}

func BenchmarkSetPrimitive(b *testing.B) {
	s := make(map[int]bool)

	for n := 0; n < b.N; n++ {
		_, ok := s[n]
		if !ok {
			s[n] = true
		}
	}
}
