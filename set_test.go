package set

import (
	"log"
	"testing"
)

func TestSetBasicUsage(t *testing.T) {
	s := New()

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
