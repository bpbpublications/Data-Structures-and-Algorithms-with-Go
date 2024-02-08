package selection

func SelectionSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		min := array[i]
		pos := i
		for j := i + 1; j < len(array); j++ {
			if array[j] < min {
				min = array[j]
				pos = j
			}
		}
		array[pos] = array[i]
		array[i] = min
	}
	return array
}
