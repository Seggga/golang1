package main

import (
	"fmt"

	"github.com/seggga/golang1/04_sort/sorter"
)

func main() {

	// неотсортированный слайс
	userSlice := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Printf("%#v", userSlice)

	// сортировка слайса
	sorter.InsSorter(userSlice)
	fmt.Printf("\n%#v", userSlice)

}
