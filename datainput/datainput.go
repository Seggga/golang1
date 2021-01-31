package datainput

import (
	"fmt"
	"strconv"
)

//IntFromStd - считывает целое число из StdIn
func IntFromStd(message string) (int, error) {

	var num int
	var userInput string

	//ввод числа
	fmt.Printf("%s ", message)
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		return 0, err
	}

	//парсинг  числа
	num, err = strconv.Atoi(userInput)
	if err != nil {
		return 0, err
	}

	return num, nil
}
