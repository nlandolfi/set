package main

import (
	"log"

	"github.com/nlandolfi/set"
)

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

/* Output:
2015/11/24 10:26:13 {3, 1, 2}
2015/11/24 10:26:13 Deck: {(10, ♥), (J, ♦), (K, ♦), (5, ♦), (6, ♦), (7, ♠), (9, ♦), (9, ♣), (9, ♠), (2, ♦), (3, ♦), (10, ♦), (8, ♣), (Q, ♥), (Q, ♦), (Q, ♣), (2, ♠), (2, ♣), (3, ♣), (4, ♥), (4, ♣), (10, ♠), (6, ♥), (8, ♥), (Q, ♠), (A, ♠), (J, ♣), (K, ♣), (5, ♠), (5, ♥), (9, ♥), (A, ♥), (A, ♣), (4, ♠), (J, ♠), (8, ♠), (8, ♦), (10, ♣), (K, ♠), (K, ♥), (6, ♠), (7, ♦), (A, ♦), (5, ♣), (6, ♣), (7, ♥), (7, ♣), (2, ♥), (3, ♠), (3, ♥), (4, ♦), (J, ♥)}
2015/11/24 10:26:13 Number of Cards: 52
2015/11/24 10:26:13 Union of ranks and suits, {J, ♥, ♦, Q, A, 5, 8, K, 2, ♠, 6, 7, 9, 3, 10, 4, ♣}
2015/11/24 10:26:13 Power set of {1, 2, 3}: {{1, 3}, {2}, {3}, {2, 3}, {}, {1}, {2, 1}, {3, 1, 2}}
*/
