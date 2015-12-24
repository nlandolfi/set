package relation_test

import (
	"testing"

	"github.com/nlandolfi/set"
	"github.com/nlandolfi/set/relation"
)

var numbers set.Interface = set.New()

func init() {
	for n := 0; n < 100; n++ {
		numbers.Add(n)
	}
}

var equality relation.AbstractInterface = relation.NewFunctionBinaryRelation(numbers, func(x, y set.Element) bool {
	return x == y
})

var lessEqual relation.AbstractInterface = relation.NewFunctionBinaryRelation(numbers, func(x, y set.Element) bool {
	return x.(int) <= y.(int)
})

func BenchmarkSymmetric(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if relation.Symmetric(equality) != true {
			b.Fatalf("The equality relation should be symmetric")
		}
	}
}

func BenchmarkTransitive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if relation.Transitive(lessEqual) != true {
			b.Fatalf("The less than or equal to relation should be transitive")
		}
	}
}
