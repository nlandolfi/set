/*
Package set defines the Interface for working with sets and
provides an implementation of a set backed by a Go map.

Furthermore we define the interface for a BinaryRelation
defined over a set. Implementations of a map backed binary
relation and predicate backed binary relation are provided.

Example:
	func main() {
		// Recall that there is no guaranteed ordering of the keys in a map,
		// therefore the structures in the following examples will change.
		s := set.New()

		s.Add(1)
		s.Add(2)
		s.Add(3)

		log.Printf("%s", s) // {1, 2, 3}

		ranks := set.With([]set.Element{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"})
		suits := set.With([]set.Element{"♠", "♥", "♦", "♣"})
		deck := set.CartesianProduct(ranks, suits)

		log.Printf("Deck: %s", deck) // {("2", "♠"), ("3", "♠"), ..., ("A", "♣") }
		log.Printf("Number of Cards: %d", deck.Cardinality()) // 52

		log.Printf("Union of ranks and suits, %v", set.Union(ranks, suits)) // {"2", "3", ..., "♦", "♣"}
		log.Printf("Power set of {1, 2, 3}: %v", set.PowerSet(s)) // {{}, {1}, {2}, {3}, {1, 2}, {1, 3}, {2, 3}, {1, 2, 3}}
	}

Note that in the majority of cases, a set of type T may be
implemented using a map[T]bool. This idiomatic solution
also avoids the extra machinery introduced by this package.
Using a map built-in type is about 3 times as fast.
*/
package set
