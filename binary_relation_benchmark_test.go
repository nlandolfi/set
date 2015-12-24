package set_test

import (
	"testing"

	"github.com/nlandolfi/set"
)

var numbers set.Interface = set.New()

func init() {
	for n := 0; n < 100; n++ {
		numbers.Add(n)
	}
}

var equality set.BinaryRelation = set.NewFunctionBinaryRelation(numbers, func(x, y set.Element) bool {
	return x == y
})

var lessEqual set.BinaryRelation = set.NewFunctionBinaryRelation(numbers, func(x, y set.Element) bool {
	return x.(int) <= y.(int)
})

func BenchmarkSymmetric(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if set.Symmetric(equality) != true {
			b.Fatalf("The equality relation should be symmetric")
		}
	}
}

func BenchmarkTransitive(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if set.Transitive(lessEqual) != true {
			b.Fatalf("The less than or equal to relation should be transitive")
		}
	}
}
