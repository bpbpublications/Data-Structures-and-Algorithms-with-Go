package seqsearch

func SeqSearch(array []int, key int) int {
	for i := 0; i < len(array); i++ {
		if key == array[i] {
			return i
		}
	}
	return -1
}

func SeqSearchRange(array []int, key int) int {
	for i, elem := range array {
		if key == elem {
			return i
		}
	}
	return -1
}
