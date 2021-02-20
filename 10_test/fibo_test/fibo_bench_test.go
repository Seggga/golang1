package fibo_test

import (
	"testing"

	"github.com/seggga/golang1/05_recursion/fibo"
)

const num = 45

func BenchmarkFindNumRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibo.FindNumRecursive(num)
	}
}

func BenchmarkFindNumRecursiveMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var optiMap = make(map[uint]uint)
		optiMap[0] = 0
		optiMap[1] = 1

		fibo.FindNumRecursiveMap(num, optiMap)
	}
}

func BenchmarkFindNumLinear(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		fibo.FindNumLinear(num)
	}
}
