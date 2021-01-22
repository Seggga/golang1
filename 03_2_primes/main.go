/*
Написать приложение, которое ищет все простые числа от 0 до N включительно.
Число N должно быть задано из стандартного потока ввода
*/

package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/seggga/golang1/03_2_primes/primes"
)

func main() {

	var num int
	var userInput string
	var primeNumbers []int = make([]int, 1)

	//приветствие
	fmt.Printf("\n\nПрограмма PRIMES находит все простые числа от 0 до N.\n\n")

	//ввод числа
	fmt.Print("Введите целое число N больше 2 (пример 10): ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("Ошибка ввода, программа завершает работу.")
		os.Exit(2)
	}

	//парсинг  числа
	num, err = strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Введенные данные не удается распознать как целое число.")
		fmt.Println("Программа завершает работу.")
		os.Exit(3)
	}

	if num <= 2 {
		fmt.Println("Введенное число должно быть больше 2. Программа завершает работу.")
		os.Exit(4)
	}

	//первое простое число
	primeNumbers[0] = 2

	for i := 3; i <= num; i++ {
		if primes.IsPrime(i, primeNumbers) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	fmt.Printf("\nВ диапазоне от 0 до %d имеются следующие простые числа:\n%v\n\n", num, primeNumbers)
}
