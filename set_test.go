package set_test

import (
	"log"
	"testing"
	"time"

	"github.com/nlandolfi/set"
)

// --- TestBasicUsage {{{

func TestBasicUsage(t *testing.T) {
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

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
	t.Parallel()

	if set.Empty.Cardinality() != 0 {
		log.Fatal("Empty set should have cardinality 0")
	}

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

// --- TestEquivalent {{{

func TestEquivalent(t *testing.T) {
	t.Parallel()

	elements := []set.Element{"A", "B", "C", "D", "E", "F"}
	A := set.WithElements(elements...)
	B := set.WithElements(elements...)

	if !set.Equivalent(A, B) {
		t.Fatalf("%s and %s should be equivalent", A, B)
	}

	if set.Equivalent(A, set.Empty) {
		t.Fatalf("%s and %s should not be equivalent", A, set.Empty)
	}

	if !set.Equivalent(set.Empty, set.Empty) {
		t.Fatalf("%s and %s should be equivalent", set.Empty, set.Empty)
	}
}

// --- }}}

// --- TestSetRelations {{{

func TestSetRelations(t *testing.T) {
	t.Parallel()

	elements := []set.Element{"A", "B", "C", "D", "E", "F"}
	A := set.WithElements(elements...)
	B := set.WithElements(elements...)

	if !set.IsSubset(A, B) {
		t.Fatalf("%s should be a subset of %s", A, B)
	}

	if set.IsProperSubset(A, B) {
		t.Fatalf("%s should be a not proper subset of %s", A, B)
	}

	if !set.IsSuperset(A, B) {
		t.Fatalf("%s should be a superset of %s", A, B)
	}

	if !set.IsSubset(A, A) {
		t.Fatalf("%s should be a subset of itself")
	}

	if set.IsProperSubset(A, A) {
		t.Fatalf("%s should not be a proper subset of itself")
	}

	if !set.IsSuperset(A, A) {
		t.Fatalf("%s should be a superset of itself")
	}

	A.Remove("A")

	if !set.IsSubset(A, B) {
		t.Fatal("%s should be a subset of %s", A, B)
	}

	if set.IsSubset(B, A) {
		t.Fatalf("%s should not be a subset of %s", B, A)
	}

	if !set.IsSuperset(B, A) {
		t.Fatal("%s should be a superset of %s", B, A)
	}

	if set.IsSuperset(A, B) {
		t.Fatal("%s should not be a superset of %s", A, B)
	}

	if !set.IsProperSubset(A, B) {
		t.Fatal("%s should be a proper subset of %s", A, B)
	}
}

// --- }}}

// --- TestSetOperations {{{

func TestSetUnion(t *testing.T) {
	t.Parallel()

	A := set.WithElements(1, 2, 3)
	B := set.WithElements(4, 5, 6)
	U := set.Union(A, B)

	for i := 1; i < 7; i++ {
		if !U.Contains(i) {
			t.Fatalf("Union of %s and %s should contain %d", A, B, i)
		}
	}

	if U.Cardinality() != 6 {
		t.Fatalf("Expected cardinality of union of %s and %s to be 6", A, B)
	}

	U.Remove(1)
	U.Remove(6)

	if !A.Contains(1) || !B.Contains(6) {
		t.Fatalf("Modifying the union set should not change the original set")
	}

}

func TestSetIntersection(t *testing.T) {
	t.Parallel()

	A := set.WithElements("one", "two")
	B := set.WithElements("two", "three")
	I := set.Intersection(A, B)

	if I.Cardinality() != 1 {
		t.Fatalf("Intersection of %s and %s should have cardinality 1", A, B)
	}

	expected := set.WithElements("two")

	if !set.Equivalent(I, expected) {
		t.Fatalf("Expected %s intersect %s to be %s, but got %s", A, B, expected, I)
	}
}

func TestSetComplement(t *testing.T) {
	t.Parallel()

	A := set.WithElements(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	B := set.WithElements(2, 4, 6, 8, 10)
	Odds := set.Complement(A, B)

	expectedOdds := set.WithElements(1, 3, 5, 7, 9)

	if !set.Equivalent(Odds, expectedOdds) {
		t.Fatalf("Expected %s\\%s to be %s, but got %s", A, B, expectedOdds, Odds)
	}
}

// --- }}}
