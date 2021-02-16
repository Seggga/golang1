package main

import (
	"fmt"
	"os"

	"github.com/seggga/golang1/04_sort/sorter"
)

func main() {

	//reading data from file
	var fileName string = "numbers.txt"
	userSlice, err := sorter.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error rading data from file\n%v\n\n", err)
	}

	//if something has been read, that is data to work on
	if userSlice == nil {
		return
	}

	//printing initial slice
	fmt.Printf("%#v", userSlice)

	// sorting and printing sorted slice
	sorter.InsSorter(userSlice)
	fmt.Printf("\n%#v", userSlice)

}
