package sorter

import (
	"bufio"
	"os"
	"strconv"
)

// InsSorter реализует алгоритм сортировки вставкой, принимает на вход слайс целых чисел []int, его же и модифицирует
func InsSorter(dataSlice []int) {
	for i := 1; i < len(dataSlice); i++ {
		for j := i; j > 0; j-- {
			if dataSlice[j] < dataSlice[j-1] {
				dataSlice[j], dataSlice[j-1] = dataSlice[j-1], dataSlice[j]
			} else {
				break
			}
		}
	}
}

// ReadFile считывает целые числа из файла, формирует выходной слайс
// требования к формату входного файла: целые числа через пробел в одну строку или несколько строк.
func ReadFile(fileName string) ([]int, error) {

	// открытие файла
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	var outSlice []int

	// получение контента из файла
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.ParseInt(line, 10, 64)

		if err != nil {
			return outSlice, err
		}

		typedNum := int(num)
		outSlice = append(outSlice, typedNum)
	}

	file.Close()

	return outSlice, err
}
