package main

import (
	"fmt"

	"github.com/seggga/golang1/04_sort/sorter"
)

func main() {

	userSlice := []int32{5, 4, 3, 2, 1, 1000, 123312, -192, 1293, 929999, 12, 3213, 423, 44, 1, 654, 5, 9}

	//printing initial slice
	fmt.Printf("%#v", userSlice)

	//printing sorted slice (insert-sort)
	fmt.Printf("\n%#v", sorter.InsSort(userSlice))

	//printing sorted slice (bubble-sort)
	fmt.Printf("\n%#v", sorter.BubbleSort(userSlice))
}
