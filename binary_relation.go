package set

// --- Types {{{

type (

	// A BinaryRelation is the interface for
	// a Binary relation from set theory.
	BinaryRelation interface {
		Universe() Interface
		ContainsRelation(Element, Element) bool
	}

	// A physical BinaryRelation is constructed
	// piecewise using the AddRelation function
	// It's representation is finite, and stored
	// completely. Contrast with a Binary Relation
	// defined by a function, f: X → ℝ
	PhysicalBinaryRelation interface {
		BinaryRelation
		AddRelation(Element, Element)
		RemoveRelation(Element, Element)
	}
)

/// --- }}}

// NewPhysicalBinaryRelationOn constructs a new
// BinaryRelation using guile's interal binaryRelation
// implementation
func NewPhysicalBinaryRelationOn(universe Interface) PhysicalBinaryRelation {
	return &binaryRelation{
		universe:  universe,
		relations: make(map[Element]map[Element]bool),
	}
}

// --- Binary Relation Implementation {{{

// binaryRelation is guile's internal representation of
// a binaryRelation
type binaryRelation struct {
	universe  Interface
	relations map[Element]map[Element]bool
}

// Universe() returns the set over which the binary
// relation is defined
func (b *binaryRelation) Universe() Interface {
	return b.universe
}

// assert is a helper function to provide
// moderate runtime type checking on the Element interface
func assert(flag bool, s string) {
	if !flag {
		panic(s)
	}
}

// AddRelation will note the fact that e1 is related to e2
// Denote our binary relation as B, then e1 B e2 <=> AddRelation(e1, e2)
func (b *binaryRelation) AddRelation(e1, e2 Element) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	var bucket map[Element]bool
	var exists bool

	// Add Normal Relation
	bucket, exists = b.relations[e1]

	if !exists {
		bucket = map[Element]bool{e2: true}
	} else {
		bucket[e2] = true
	}

	b.relations[e1] = bucket
}

// RemoveRelation is the inverse operation of AddRelation
// It works regardless of whether the relation is actually present
func (b *binaryRelation) RemoveRelation(e1, e2 Element) {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	if bucket, exists := b.relations[e1]; exists {
		if _, exists := bucket[e2]; exists {
			delete(bucket, e2)
		}
	}
}

// ContainsRelation determines whether the given relation exists and is
// defined as a member of this binary relation. Note: Order of e1, and e2
// matters, of course.
func (b *binaryRelation) ContainsRelation(e1, e2 Element) bool {
	assert(b.universe.Contains(e1), "(*binaryRelation).AddRelation: element 1 is not contained in universe")
	assert(b.universe.Contains(e2), "(*binaryRelation).AddRelation: element 2 is not contained in universe")

	if bucket, exists := b.relations[e1]; exists {
		if _, defined := bucket[e2]; defined {
			return true
		}
	}

	return false
}

// --- }}}

// --- Properties {{{

// Reflexive checks the following condition:
//  xBx for any x ∈ X ≡ Universe()
func Reflexive(b BinaryRelation) bool {
	for _, e := range b.Universe().Elements() {
		if !b.ContainsRelation(e, e) {
			return false
		}
	}

	return true
}

// Complete checks the following condition:
//  xBy or yBx for any x, y ∈ X ≡ Universe()
func Complete(b BinaryRelation) bool {
	elems := b.Universe().Elements()

	// n^2! yuck!
	for _, x := range elems {
		for _, y := range elems {
			if !(b.ContainsRelation(x, y) || b.ContainsRelation(y, x)) {
				return false
			}
		}
	}

	return true
}

// Transitive checks the following condition:
//	 (xBy and yBz) ⇒  xBz for any x, y, z ∈ X ≡ Universe()
func Transitive(b BinaryRelation) bool {
	if !Complete(b) {
		return false
	}

	elems := b.Universe().Elements()

	// n^3 :(
	for _, x := range elems {
		for _, y := range elems {
			for _, z := range elems {
				if b.ContainsRelation(x, y) && b.ContainsRelation(y, z) {
					if !b.ContainsRelation(x, z) {
						return false
					}
				}
			}
		}
	}

	return true
}

// Symmetric checks the following condition:
//	 xBy ⇒  yBx for any x, y ∈ X ≡ Universe
func Symmetric(b BinaryRelation) bool {
	elems := b.Universe().Elements()

	for _, x := range elems {
		for _, y := range elems {
			if b.ContainsRelation(x, y) {
				if !b.ContainsRelation(x, y) {
					return false
				}
			}
		}
	}

	return true
}

// AntiSymmetric checks the following condition:
//	(xBy and yBx) ⇒  (x = y), for any x, y ∈ X
func AntiSymmetric(b BinaryRelation) bool {
	elems := b.Universe().Elements()
	for _, x := range elems {
		for _, y := range elems {
			if b.ContainsRelation(x, y) && b.ContainsRelation(y, x) {
				if x != y {
					return false
				}
			}
		}
	}

	return true
}

func ComposableRelations(relations []BinaryRelation) bool {
	if len(relations) == 0 {
		return true
	}

	u := relations[0].Universe()

	for _, b := range relations {
		if !Equivalent(u, b.Universe()) {
			return false
		}
	}

	return true
}

// --- }}}

// --- Orders {{{

// WeakOrder checks the B is Complete and Transitive
// i.e, > and >= defined on the universe of naturals
func WeakOrder(b BinaryRelation) bool {
	return Complete(b) && Transitive(b)
}

// StrictOrder checks that B is a weak order and additionally
// that B is Symmetric. (This is > verse >=)
func StrictOrder(b BinaryRelation) bool {
	return WeakOrder(b) && AntiSymmetric(b)
}

func Reverse(b BinaryRelation) BinaryRelation {
	return NewFunctionBinaryRelation(b.Universe(), func(x, y Element) bool {
		return !b.ContainsRelation(x, y)
	})
}

// --- }}}

// --- Function Based Binary Relation {{{

type RelatedFunction func(Element, Element) bool

type fnBinaryRelation struct {
	universe Interface
	related  RelatedFunction
}

func NewFunctionBinaryRelation(u Interface, fn RelatedFunction) BinaryRelation {
	return &fnBinaryRelation{
		universe: u,
		related:  fn,
	}
}

func (fb *fnBinaryRelation) Universe() Interface {
	return fb.universe
}

func (fb *fnBinaryRelation) ContainsRelation(x, y Element) bool {
	return fb.related(x, y)
}

// --- }}}
