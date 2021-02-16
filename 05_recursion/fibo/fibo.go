package fibo

//FindNumRecursive - рекурсивно вычисляет n-ый элемент последовательности Фибоначчи
func FindNumRecursive(n uint) uint {

	switch {
	case n == 0:
		return 0
	case n == 1:
		return 1
	default:
		return FindNumRecursive(n-2) + FindNumRecursive(n-1)
	}
}

//FindNumLinear - выполняет линейные вычисления n-го числа Фибоначчи (без рекурсии)
func FindNumLinear(n uint) uint {

	var fiboMinus2 uint = 0 // нулевой член ряда
	var fiboMinus1 uint = 1 // первый член ряда
	var i uint

	for i = 2; i <= n; i++ {
		fiboMinus2, fiboMinus1 = fiboMinus1, fiboMinus1+fiboMinus2
	}
	return fiboMinus1
}

//FindNumRecursiveMap - выполняет рекурсивное вычисление n-го члена ряда Фибоначчи
func FindNumRecursiveMap(n uint, optiMap map[uint]uint) uint {

	result, ok := optiMap[n]
	if ok == true {
		return result
	}

	optiMap[n] = FindNumRecursiveMap(n-2, optiMap) + FindNumRecursiveMap(n-1, optiMap)
	return optiMap[n]
}
