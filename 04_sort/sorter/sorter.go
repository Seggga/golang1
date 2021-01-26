package sorter

// InsSorter реализует алгоритм сортировки вставкой, принимает на вход слайс целых чисел []int, его же и модифицирует
func InsSorter(dataSlice []int) {
	for i := 1; i < len(dataSlice); i++ {
		for j := i; j > 0; j-- {
			if dataSlice[j] < dataSlice[j - 1] {
				dataSlice[j], dataSlice[j - 1] = dataSlice[j - 1], dataSlice[j]
			} else {
				break
			}
		}		
	}
}