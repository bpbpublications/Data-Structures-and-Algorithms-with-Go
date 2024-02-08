package main

import (
	"fmt"
	"sort"
	"sorting/bubble"
	"sorting/insertion"
	"sorting/quick"
	"sorting/selection"
)

func main() {
	array := []int{18, 1, 5, 27, 8, 25, 9, 21}
	fmt.Println(insertion.InsertionSort(array))

	array = []int{18, 1, 5, 27, 8, 25, 9, 21}
	fmt.Println(selection.SelectionSort(array))

	array = []int{18, 1, 5, 27, 8, 25, 9, 21}
	fmt.Println(bubble.BubbleSort(array))

	array = []int{18, 1, 5, 27, 8, 25, 9, 21}
	fmt.Println(quick.QuickSort(array, 0, len(array)-1))

	array = []int{18, 1, 5, 27, 8, 25, 9, 21}
	sort.Ints(array)
	fmt.Println(array)
}
