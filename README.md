## set [![GoDoc](https://godoc.org/github.com/nlandolfi/set?status.svg)](https://godoc.org/github.com/nlandolfi/set)

The set package defines the Interface for working with sets and provides an implementation of a set backed by a Go map.

Furthermore we define the interface for a BinaryRelation defined over a set. Implementations of a map backed binary relation and predicate backed binary relation are provided.

Note: In the majority of cases, a set of type `T` can be represented as a `map[T]bool`, without the excess machinery. Using a go map is about 3 times as fast.
