package set_test

import (
	"log"
	"testing"

	"github.com/nlandolfi/set"
)

// --- TestBasicUsage {{{

func TestBasicUsage(t *testing.T) {
	s := set.New()

	s.Add(1)
	s.Add(2)
	s.Add(3)

	if !s.Contains(1) {
		t.Errorf("Expected set to contain 1")
	}

	if !s.Contains(2) {
		t.Errorf("Expected set to contain 2")
	}

	if !s.Contains(3) {
		t.Errorf("Expected set to contain 3")
	}

	if s.Contains(4) {
		t.Errorf("Expected set to not contain 4")
	}

	if s.Cardinality() != 3 {
		t.Errorf("Expected set cardinality to be 3")
	}

	s.Remove(2)

	if s.Contains(2) {
		t.Errorf("Expected set to no longer contain 2")
	}

	if len(s.Elements()) != int(s.Cardinality()) {
		log.Printf("Elements %+v", s.Elements())
		log.Printf("s: %+v", s)
		t.Errorf("Expected length of Elements() [%d] to match Cardinality() [%d]", len(s.Elements()), s.Cardinality())
	}
}

// --- }}}

// --- TestConstructors {{{

func TestConstructors(t *testing.T) {
	s := make([]set.Element, 10)
	for i := 0; i < 10; i++ {
		s[i] = i
	}

	A := set.With(s)
	B := set.WithElements(s...)

	for _, v := range s {
		if !A.Contains(v) {
			log.Fatal("A should contain %d (created by set.With)", v)
		}
		if !B.Contains(v) {
			log.Fatal("B should contain %d (created with set.WithElements", v)
		}
	}
}

// --- }}}

// --- TestMembership {{{

func TestMembership(t *testing.T) {
	testElements := make(map[int]bool)
	for i := 0; i < 123; i++ {
		testElements[i] = true
	}

	s := set.New()

	for k, _ := range testElements {
		if s.Contains(k) {
			log.Fatalf("Set should not contain: %d", k)
		}

		s.Add(k)
	}

	for k, _ := range testElements {
		if !s.Contains(k) {
			log.Fatalf("Set should contain %d", k)
		}

		s.Remove(k)
	}

	for k, _ := range testElements {
		if s.Contains(k) {
			log.Fatal("Set should not contain %d", k)
		}
	}
}

// --- }}}
