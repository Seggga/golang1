package main

import (
	"fmt"
	//"sorter"
)

func main() {

	userSlice := []int{9,8,7,6,5,4,3,2,1}
	fmt.Printf("%#v", userSlice)

	sorter.InsSorter(userSlice)
	fmt.Printf("\n%#v", userSlice)

}