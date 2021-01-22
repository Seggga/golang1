/*
 Доработать калькулятор: больше операций и валидации данных.

*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	var a, b, result float64
	var userInput string

	//Приветствие
	fmt.Printf("\n\nПрограмма КАЛЬКУЛЯТОР вычисляет значения простейших арифметических выражений.\n\n")

	//ввод первого числа
	fmt.Print("Введите первое число (пример 12.345): ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Ошибка ввода, программа завершает работу.")
		os.Exit(2)
	}
	//парсинг первого числа
	a, err = strconv.ParseFloat(userInput, 32)
	if err != nil {
		fmt.Println("Введенные данные не удается распознать как число.")
		fmt.Println("Программа завершает работу.")
		os.Exit(3)
	}

	//ввод второго числа
	fmt.Print("Введите второе число (пример 0.987): ")
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Ошибка ввода, программа завершает работу.")
		os.Exit(2)
	}
	//парсинг второго числа
	b, err = strconv.ParseFloat(userInput, 32)
	if err != nil {
		fmt.Println("Введенные данные не удается распознать как число.")
		fmt.Println("Программа завершает работу.")
		os.Exit(3)
	}

	//ввод знака арифметического действия
	fmt.Print("Введите знак арифметического действия ( + - * / ): ")
	_, err = fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Ошибка ввода, программа завершает работу.")
		os.Exit(2)
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
		fmt.Println("Введенные данные не удается распознать как знак арфиметического действия.")
		fmt.Println("Программа завершает работу.")
		os.Exit(4)
	}

	fmt.Printf("\nЗначение выражения %f %s %f равно %f\n\n\n", a, userInput, b, result)

}
