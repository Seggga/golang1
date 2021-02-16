package main

import (
	"fmt"
	"os"

	"github.com/seggga/golang1/05_recursion/fibo"
	"github.com/seggga/golang1/datainput"
)

func main() {

	message := `Программа вычисляет число в ряде Фибоначчи (0 1 1 2 3 5 ...).
Введите номер позиции для поиска (целое число больше 2):`
	num, err := datainput.IntFromStd(message)
	if err != nil {
		fmt.Printf("Ошибка ввода, %v", err)
		os.Exit(2)
	}

	if num < 2 {
		fmt.Println("Введенное число меньше 2. Программа завершает работу.")
		os.Exit(3)
	}

	n := uint(num)
	// расчет с использованием линейной функции
	fmt.Printf("position %d : number %d\n", n, fibo.FindNumLinear(n))

	// расчет с использованием рекурсивной функции
	fmt.Printf("position %d : number %d\n", n, fibo.FindNumRecursive(n))

	// расчет с использованием оптимизированной рекурсивной функции
	optiMap := make(map[uint]uint)
	optiMap[0] = 0
	optiMap[1] = 1
	fmt.Printf("position %d : number %d\n", n, fibo.FindNumRecursiveMap(n, optiMap))

}
