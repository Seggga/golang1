package fibo

import "testing"
	
const num = 45

func BenchmarkFindNumRecursive(b *testing.B ) {
	for i := 0; i < b.N; i++ {
		FindNumRecursive(num)
	}
}

func BenchmarkFindNumRecursiveMap(b *testing.B) {
	for i := 0; i < b.N; i++ {

		var  optiMap = make(map[uint]uint)
		optiMap[0] = 0
		optiMap[1] = 1

		FindNumRecursiveMap(num, optiMap)
	}
}

func BenchmarkFindNumLinear(b *testing.B ) {
	for i := 0; i <= b.N; i++ {
		FindNumLinear(num)
	}
}
