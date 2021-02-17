package sorter

// InsSort реализует алгоритм сортировки вставкой, принимает на вход слайс целых чисел []int, его же и модифицирует
func InsSort(userSlice []int32) []int32 {

	sortedSlice := make([]int32, len(userSlice))
	copy(sortedSlice, userSlice)

	for i := 1; i < len(sortedSlice); i++ {
		for j := i; j > 0; j-- {
			if sortedSlice[j] < sortedSlice[j-1] {
				sortedSlice[j], sortedSlice[j-1] = sortedSlice[j-1], sortedSlice[j]
			} else {
				break
			}
		}
	}

	return sortedSlice
}

//BubbleSort implements sorting algorythm on a slice []int32
func BubbleSort(userSlice []int32) []int32 {

	sortedSlice := make([]int32, len(userSlice))
	copy(sortedSlice, userSlice)

	for i := 0; i < len(sortedSlice)-1; i++ {

		for j := 0; j < len(sortedSlice)-i-1; j++ {
			if sortedSlice[j] > sortedSlice[j+1] {
				sortedSlice[j], sortedSlice[j+1] = sortedSlice[j+1], sortedSlice[j]
			}
		}
	}

	return sortedSlice
}
