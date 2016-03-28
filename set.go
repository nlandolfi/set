package set

import (
	"fmt"
	"strings"
)

// --- Types {{{

type (
	// An Element can be any type. The types of various items added
	// to a set likely should not change, though this is not strictly
	// disallowed.
	Element interface{}

	// Elements is a typed slice of elements
	Elements []Element

	// AbstractInterface requires that a structure can determine membership
	//
	// Note: One could declare predicate-based membership, which is the
	// preferred method of implementing infinite sets.
	AbstractInterface interface {
		// Contains returns a boolean indicating set membership.
		// True iff e ∈ the set backing this AbstractInterface.
		Contains(e Element) bool
	}

	// Interface includes the basic operations which may be performed on a
	// physical set.
	Interface interface {
		// Inherits: Contains(Element) bool
		AbstractInterface

		// Add includes e as a member of the set.
		//
		// Note: Add must be idempotent.
		Add(e Element)

		// Remove excludes e as a member of the set.
		//
		// Note: Remove must be idempotent.
		Remove(e Element)

		// Cardinality retrieves the size of the set.
		Cardinality() uint

		// Elements retrieves a slice of all member elements.
		//
		// Note: Mutating this slice must not modify the
		// underlying set structure.
		Elements() []Element
	}
)

// --- }}}

// --- Constructors {{{

// New constructs a Set object
func New() Interface {
	return &mapSet{}
}

// With constructs a new Set object,
// containing the elements in the slice, 'elements'
func With(elements []Element) Interface {
	s := New()

	for i := range elements {
		s.Add(elements[i])
	}

	return s
}

// WithElements constructs a set with variable number of elements.
func WithElements(elements ...Element) Interface {
	return With(elements)
}

// --- }}}

// --- MapSet {{{

// mapSet is an implementation of a set backed by a golang map
type mapSet map[Element]bool

// Add includes e as a member of the set.
//
// Add is idempotent.
func (s *mapSet) Add(e Element) {
	(*s)[e] = true
}

// Remove excludes e as a member of the set.
//
// Remove is idempotent.
func (s *mapSet) Remove(e Element) {
	delete(*s, e)
}

// Contains returns a flag determining whether an Element, e
// is a member of the set
func (s *mapSet) Contains(e Element) bool {
	_, contains := (*s)[e]

	// If e is a set, check deep equality with keys
	sub, ok := e.(Interface)
	if ok {
		for k := range *s {
			key, ok := k.(Interface)
			if !ok {
				continue
			}

			if Equivalent(sub, key) {
				contains = true
				break
			}
		}
	}

	return contains
}

// Cardinality returns the size of the set.
// Cardinality(s) ≡ |s|
func (s *mapSet) Cardinality() uint {
	return uint(len(*s))
}

// Elements returns a slice of the elements contained in this
// set.
//
// Note: This slice is not the internal reprentation and therefore
// can be mutated.
func (s *mapSet) Elements() []Element {
	e := make([]Element, len(*s))
	i := 0
	for k := range *s {
		e[i] = k
		i++
	}
	return e
}

// Returns a stream of the set elements
func (s *mapSet) Iter() <-chan Element {
	c := make(chan Element, len(*s))

	els := (*s).Elements()

	go func() {
		for i := range els {
			c <- els[i]
		}

		close(c)
	}()

	return c
}

func (s *mapSet) String() string {
	return String(s)
}

// --- }}}

// --- Tuple (For Cartesian Products) {{{

// Tuple represents a two-dimensional list of Elements.
type Tuple struct {
	First, Second Element
}

// String constructs a string representation of a Tuple.
// For Example: (First, Second).
func (t *Tuple) String() string {
	return fmt.Sprintf("(%v, %v)", t.First, t.Second)
}

// --- }}}

// --- Equivalence, IsSubset IsSuperset {{{

// Equivalent → true iff s1 ≡ s2 (s1 is identical to s2)
func Equivalent(s1, s2 Interface) bool {
	// is every element in s1 a member of s2
	for _, e := range s1.Elements() {
		if !s2.Contains(e) {
			return false
		}
	}

	// is every element in s2 a member of s1
	for _, e := range s2.Elements() {
		if !s1.Contains(e) {
			return false
		}
	}

	return true
}

// IsSubset → true iff s1 ⊆ s2 (s1 is a subset of s2)
func IsSubset(s1, s2 Interface) bool {
	for _, e := range s1.Elements() {
		if !s2.Contains(e) {
			return false
		}
	}

	return true
}

// IsProperSubset → true iff s1 ⊊ s2 (s1 is the proper subset of s2)
func IsProperSubset(s1, s2 Interface) bool {
	return IsSubset(s1, s2) && !Equivalent(s1, s2)
}

// IsSuperset → true iff s2 ⊆ s1 (s1 is a superset of s2).
func IsSuperset(s1, s2 Interface) bool {
	return IsSubset(s2, s1)
}

// --- }}}

// --- Union, Intersection {{{

// Union → s1 ∪ s2
func Union(s1, s2 Interface) Interface {
	s := With(s1.Elements())

	for _, e := range s2.Elements() {
		s.Add(e)
	}

	return s
}

// Intersection → s1 ∩ s2
func Intersection(s1, s2 Interface) Interface {
	s := New()

	c1, c2 := s1.Cardinality(), s2.Cardinality()

	if c1 < c2 {
		for _, e := range s1.Elements() {
			if s2.Contains(e) {
				s.Add(e)
			}
		}
	} else {
		for _, e := range s2.Elements() {
			if s1.Contains(e) {
				s.Add(e)
			}
		}
	}

	return s
}

// Complement →  s1\s2 (the relative complement of s2 with s1)
// That is, all elements in s1 that are not in s2.
func Complement(s1, s2 Interface) Interface {
	s := New()

	for _, e := range s1.Elements() {
		if !s2.Contains(e) {
			s.Add(e)
		}
	}

	return s
}

// --- }}}

// --- Misc. {{{

// Empty is the empty set, ∅
var Empty = New()

// String generates a string representation of a set of the form
// "{element1, element2, ..., elementN}".
func String(s Interface) string {
	elements := s.Elements()

	elementStrings := make([]string, len(elements))

	for i := range s.Elements() {
		elementStrings[i] = fmt.Sprintf("%v", elements[i])
	}

	return fmt.Sprintf("{%s}", strings.Join(elementStrings, ", "))
}

// Clone creates a carbon copy of s1.
func Clone(s1 Interface) Interface {
	return With(s1.Elements())
}

// CartesianProduct → {(x, y) | ∀ x ∈ s1, ∀ y ∈ s2}
// For example:
//		CartesianProduct(A, B), where A = {1, 2} and B = {7, 8}
//        => {(1, 7), (1, 8), (2, 7), (2, 8)}
func CartesianProduct(s1, s2 Interface) Interface {
	s := New()

	for _, e1 := range s1.Elements() {
		for _, e2 := range s2.Elements() {
			s.Add(Tuple{First: e1, Second: e2})
		}
	}

	return s
}

// multiUnion(e, T) := {X ∪ {e} | X ∈ T}.
func multiUnion(e Element, T Interface) Interface {
	n := New()

	for _, t := range T.Elements() {
		n.Add(Union(WithElements(e), t.(Interface)))
	}

	return n
}

// PowerSet → 𝒫(s)
// Source: Implemeneted using the recursive algorithm given on wikipedia,
//   https://en.wikipedia.org/wiki/Power_set#Algorithms
func PowerSet(s Interface) Interface {
	// If S = { }, then P(S) = { { } } is returned.
	if s.Cardinality() == 0 {
		return WithElements(New())
	}

	e := s.Elements()[0]
	T := Complement(s, WithElements(e))

	return Union(PowerSet(T), multiUnion(e, PowerSet(T)))
}

// 𝒫 is an alias for the PowerSet function.
func 𝒫(s Interface) Interface {
	return PowerSet(s)
}

// --- }}}
