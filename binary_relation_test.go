package set_test

import (
	"testing"

	"github.com/nlandolfi/set"
)

func TestBinaryRelationBasicUsage(t *testing.T) {
	s := set.WithElements(1, 2, 3, 4, 5, 6, 7, 8, 9, 0)

	b := set.NewPhysicalBinaryRelationOn(s)

	if b.Universe() != s {
		t.Errorf("Expected b's universe to be the set given in constructoon")
	}

	// Let's define the greater than binary relation
	b.AddRelation(2, 1)
	b.AddRelation(3, 2)
	b.AddRelation(4, 3)
	b.AddRelation(5, 4)
	b.AddRelation(6, 5)
	b.AddRelation(7, 6)
	b.AddRelation(8, 7)
	b.AddRelation(9, 8)
	b.AddRelation(1, 0)

	if !b.ContainsRelation(3, 2) {
		t.Error("Expected the binary relation to contain (3, 2), as we defined it")
	}

	if set.Complete(b) {
		t.Error("Expected b to be incomplete, consider: (9, 0)")
	}

	if !b.ContainsRelation(1, 0) {
		t.Errorf("Expected binary relation to contain (1, 0")
	}

	b.RemoveRelation(1, 0)

	if b.ContainsRelation(1, 0) {
		t.Errorf("Expected binary relation to no longer contain (1, 0), as we removed it")
	}
}
