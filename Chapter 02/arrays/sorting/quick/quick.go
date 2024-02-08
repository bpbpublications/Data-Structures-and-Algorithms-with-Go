package quick

func QuickSort(array []int, low, high int) []int {
	if low < high {
		array, j := partition(array, low, high)
		QuickSort(array, low, j-1)
		QuickSort(array, j+1, high)
	}
	return array
}

func partition(array []int, low, high int) ([]int, int) {
	pivot := array[low]
	i := low
	j := high
	for i < j {
		for array[i] <= pivot && i < j {
			i++
		}
		for array[j] > pivot {
			j--
		}
		if i < j {
			array[i], array[j] = array[j], array[i]
		}
	}

	array[low], array[j] = array[j], pivot

	return array, j
}
