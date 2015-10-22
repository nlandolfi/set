package set

import (
	"fmt"
	"strings"
)

// --- Types {{{

type (
	// An Element can be any type. The types of various items added
	// to a set should probably not change, though this isn't strictly
	// disallowed
	Element interface{}

	// A slice of elements
	Elements []Element

	// AbstractInterface defines an abstraction over a physical set.
	// It is a set with potentially infinitie membership, (predicate definitions)
	AbstractInterface interface {
		Contains(Element) bool
	}

	// Interface represents the set of elements You can add, remove and count the number
	// of elements of a set. Additionally, you can ask for the slice of elements which
	// the set contains
	Interface interface {
		AbstractInterface

		Add(Element) bool
		Remove(Element) bool
		Cardinality() uint

		Elements() []Element
		Iter() <-chan Element
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

func WithElements(elements ...Element) Interface {
	return With(elements)
}

// --- }}}

// --- MapSet {{{

// mapSet is an implementation of a set backed by a golang map
type mapSet map[Element]bool

// Add will include Element, e, as a member of the set.
// If e is already a member of the set Add still works.
// Add returns a boolean. If the Element, e, was already
// a member of the Set Add returns true, else it is false
func (s *mapSet) Add(e Element) bool {
	_, contains := (*s)[e]

	if !contains {
		(*s)[e] = true
	}

	return contains
}

// Remove will exclude an Element, e, as a member of the set.
// If e is not a member of the set, Remove still works, but it
// will return false. If e was a member which was removed,
// Remove will return true.
func (s *mapSet) Remove(e Element) bool {
	_, contains := (*s)[e]

	if contains {
		delete(*s, e)
	}

	return contains
}

// Contains returns a flag determining whether an Element, e
// is a member of the set
func (s *mapSet) Contains(e Element) bool {
	_, ok := (*s)[e]
	return ok
}

// Cardinality returns the size of the set.
// Suppose set S, Cardinality(S) ≡ |S| (as expected)
func (s *mapSet) Cardinality() uint {
	return uint(len(*s))
}

// Elements returns a slice of the elements contained in this
// set. This slice is not the internal reprentation and therefore
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

func (s *mapSet) Iter() <-chan Element {
	c := make(chan Element, len(*s))

	go func() {
		for k := range *s {
			c <- k
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

type Tuple struct {
	First, Second Element
}

func (t *Tuple) String() string {
	return fmt.Sprintf("(%v, %v)", t.First, t.Second)
}

// --- }}}

// --- Equivalence, IsSubset IsSuperset {{{

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

func IsSubset(s1, s2 Interface) bool {
	for _, e := range s1.Elements() {
		if !s2.Contains(e) {
			return false
		}
	}

	return true
}

func IsSuperset(s1, s2 Interface) bool {
	return IsSubset(s2, s1)
}

// --- }}}

// --- Union, Intersection {{{

func Union(s1, s2 Interface) Interface {
	s := With(s1.Elements())

	for _, e := range s2.Elements() {
		s.Add(e)
	}

	return s
}

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

func String(s Interface) string {
	elements := s.Elements()

	elementStrings := make([]string, len(elements))

	for i := range s.Elements() {
		elementStrings[i] = fmt.Sprintf("%v", elements[i])
	}

	return fmt.Sprintf("{%s}", strings.Join(elementStrings, ", "))
}

func Clone(s1 Interface) Interface {
	clone := New()

	for _, e := range s1.Elements() {
		clone.Add(e)
	}

	return clone
}

func CartesianProduct(s1, s2 Interface) Interface {
	s := New()

	for _, e1 := range s1.Elements() {
		for _, e2 := range s2.Elements() {
			s.Add(&Tuple{First: e1, Second: e2})
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

func PowerSet(s Interface) Interface {
	// If S = { }, then P(S) = { { } } is returned.
	if s.Cardinality() == 0 {
		return WithElements(New())
	}

	e := s.Elements()[0]
	T := Complement(s, WithElements(e))

	return Union(PowerSet(T), multiUnion(e, PowerSet(T)))
}

// --- }}}
