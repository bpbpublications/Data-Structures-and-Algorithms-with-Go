package insertion

func InsertionSort(array []int) []int {
	for i := 1; i < len(array); i++ {
		k := array[i]
		j := i - 1
		for j >= 0 && array[j] > k {
			array[j+1] = array[j]
			j = j - 1
		}
		array[j+1] = k
	}
	return array
}
