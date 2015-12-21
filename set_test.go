package set_test

import (
	"log"
	"testing"
	"time"

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

// --- TestCardinality {{{

func TestCardinality(t *testing.T) {
	s := set.New()
	for i := 0; i < 100; i++ {
		s.Add(i)
	}

	if s.Cardinality() != 100 {
		log.Fatal("S should have cardinality 100, but has cardinality %d", s.Cardinality())
	}

	elements := s.Elements()

	if len(elements) != 100 {
		log.Fatal("S should return an []Element of length 100")
	}

	elements[46] = nil

	if !s.Contains(46) {
		log.Fatal("Mutating the elements array returned from set.Elements() should not affect set membership")
	}

	if s.Cardinality() != 100 {
		log.Fatal("Mutating the elements array returned from set.Elements() should not affect set cardinality")
	}

	count := 100

	for range s.Iter() {
		count--
	}

	if count != 0 {
		log.Fatal("set.Iter() should return 100 items, but only got %d", 100-count)
	}

	counts := 1000
	counts_channel := make(chan int)

	for i := 0; i < 10; i++ {
		go func() {
			for range s.Iter() {
				counts_channel <- 1
			}
		}()
	}

Counting:
	for {
		select {
		case <-counts_channel:
			counts--
		case <-time.After(10 * time.Millisecond):
			break Counting
		}
	}

	if counts != 0 {
		t.Fatal("Should be able to read a set's contents from multiple threads at once.")
	}
}

// --- }}}
