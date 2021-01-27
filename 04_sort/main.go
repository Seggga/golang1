package main

import (
	"fmt"
	"os"

	"github.com/seggga/golang1/04_sort/sorter"
)

func main() {

	//ввод исходной последовательности чисел из файла
	var fileName string = "numbers.txt"
	userSlice, err := sorter.ReadFile(fileName)
	if err != nil {
		fmt.Printf("ошибка считывания данных из файла\n%v\n%#v\n", err, err)
	}

	//если хоть что-то удалось считать, будем работать с этим
	if userSlice == nil {
		os.Exit(2)
	}

	//вывод исходной последовательности
	fmt.Printf("%#v", userSlice)

	// сортировка и вывод слайса
	sorter.InsSorter(userSlice)
	fmt.Printf("\n%#v", userSlice)

}
