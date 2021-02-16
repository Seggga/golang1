package main

/*
 Доработать калькулятор: больше операций и валидации данных.

*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var a, b, result float64
	var userInput string

	//Приветствие
	fmt.Println("This application implements a set of 4 simple mathematics operations.")

	//ввод первого числа
	fmt.Print("Please, enter the first operand (for example, 12.345): ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Printf("There is an error entering data.\n%v", err)
		return
	}
	
	//парсинг первого числа
	a, err = strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("Entered data cannot be recognized as a number.")
		fmt.Printf("%v", err)
		return
	}

	//ввод второго числа
	fmt.Print("Please, enter the second operand (for example, 0.987): ")
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		fmt.Printf("There is an error entering data.\n%v", err)
		return
	}
	//парсинг второго числа
	b, err = strconv.ParseFloat(userInput, 64)
	if err != nil {
		fmt.Println("Entered data cannot be recognized as a number.")
		fmt.Printf("%v", err)
		return
	}

	//ввод знака арифметического действия
	fmt.Print("Please, enter the sign of desired operation ( + - * / ): ")
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		fmt.Printf("There is an error entering data.\n%v", err)
		return
	}
	switch userInput {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		fmt.Println("Entered data cannot be recognized as a mathematics sign.")
		fmt.Printf("%v", err)
		return
	}

	fmt.Printf("\nMathematical equasion %f %s %f results in %f\n", a, userInput, b, result)

}
