package sorter_test

import (
	"testing"

	"github.com/seggga/golang1/04_sort/sorter"
)

var inSlice = []int32{1, 5, 200, 3, -15, 88, 99, 988, 5, 5, 63, 12, 0, 155, 88, 5, 32, 11, 2, 565, 8, 565, 45, 5, 321, 5, 6, 88, 45, 654, 6}

func BenchmarkInsSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sorter.InsSort(inSlice)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sorter.BubbleSort(inSlice)
	}
}
