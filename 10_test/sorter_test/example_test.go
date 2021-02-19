package sorter_test

import (
	"fmt"

	"github.com/seggga/golang1/04_sort/sorter"
)

func ExampleInsSort() {
	inSlice := []int32{5, 4, 3, 2, 1, 0}
	fmt.Printf("%v", sorter.InsSort(inSlice))
	// Output: [0 1 2 3 4 5]
}

func ExampleBubbleSort() {
	inSlice := []int32{3, 5, 0, 2, 1, 4}
	fmt.Printf("%v", sorter.BubbleSort(inSlice))
	// Output: [0 1 2 3 4 5]
}
