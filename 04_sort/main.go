package main

import (
	"fmt"

	"github.com/seggga/golang1/04_sort/sorter"
)

func main() {

	//initial slice
	userSlice := []int32{5, 4, 3, 2, 1, 1000, 123312, -192, 1293, 929999, 12, 3213, 423, 44, 1, 654, 5, 9}

	//printing initial slice
	fmt.Printf("initial slice :     %#v", userSlice)

	//printing sorted slice (insert-sort)
	fmt.Printf("\nInsSort output :    %#v", sorter.InsSort(userSlice))

	//printing sorted slice (bubble-sort)
	fmt.Printf("\nBubbleSort output : %#v", sorter.BubbleSort(userSlice))

}
