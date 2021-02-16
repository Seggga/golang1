package main

/*
Написать приложение, которое ищет все простые числа от 0 до N включительно.
Число N должно быть задано из стандартного потока ввода
*/

import (
	"fmt"
	"os"
	"strconv"

	"github.com/seggga/golang1/03_2_primes/primes"
)

func main() {

	//приветствие
	fmt.Printf("\n\nThis application looks for a set of prime numbers from 0 to N.\n\n")

	//ввод числа
	var userInput string
	fmt.Print("Please, enter a N (the limit) greater then 2 (for example, 10): ")
	_, err := fmt.Scanln(&userInput)
	if err != nil {
		fmt.Println("There is an error in entering data.")
		fmt.Printf("%v", err)
		return
	}

	//парсинг  числа
	var num int
	num, err = strconv.Atoi(userInput)
	if err != nil {
		fmt.Println("Entered data cannot be recognized as a number.")
		fmt.Printf("%v", err)
		return
	}

	if num <= 2 {
		fmt.Println("limit number ( N ) is expected to be greater then 2. Entered data is %d", num)
		return
	}

	
	var primeNumbers []int = make([]int, 1)
	primeNumbers[0] = 2  //первое простое число

	for i := 3; i <= num; i++ {
		if primes.IsPrime(i, primeNumbers) {
			primeNumbers = append(primeNumbers, i)
		}
	}

	fmt.Printf("\nIn the given interval from 0 to %d there are primes as folows:\n%v\n", num, primeNumbers)
}
