/*
Example:
	func main() {
		s := set.New()

		s.Add(1)
		s.Add(2)
		s.Add(3)

		log.Printf("%s", s)

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
