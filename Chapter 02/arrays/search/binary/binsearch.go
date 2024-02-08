package binary

func BinSearch(array []int, key int) int {
	low := 0
	high := len(array) - 1

	for low <= high {
		mid := (low + high) / 2

		if key == array[mid] {
			return mid
		} else if key < array[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
