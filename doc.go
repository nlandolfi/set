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

		log.Printf("%s", s) // => {1, 2, 3}

		ranks := set.With([]set.Element{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"})
		suits := set.With([]set.Element{"♠", "♥", "♦", "♣"})
		deck := set.CartesianProduct(ranks, suits)

		log.Printf("Deck: %s", deck)
		log.Printf("Number of Cards: %d", deck.Cardinality())

		log.Printf("Union of ranks and suits, %v", set.Union(ranks, suits))

		log.Printf("Power set of {1, 2, 3}: %v", set.PowerSet(s))
	}
*/
package set
