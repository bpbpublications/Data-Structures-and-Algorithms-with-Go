package main

import (
	"fmt"
	binary "search/binary"
	seqsearch "search/sequential"
	"sort"
)

func main() {
	array := []int{1, 18, 5, 27, 8, 25, 9, 21}
	// Sequential search
	fmt.Println(seqsearch.SeqSearch(array, 27))
	fmt.Println(seqsearch.SeqSearchRange(array, 27))

	array = []int{1, 5, 8, 9, 18, 21, 25, 27}
	// Binary search
	fmt.Println(binary.BinSearch(array, 18))

	fmt.Println(sort.SearchInts(array, 9))
}
