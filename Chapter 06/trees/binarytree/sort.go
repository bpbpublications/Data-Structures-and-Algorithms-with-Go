package binarytree

import (
	"container/heap"
)

func ArrayToTree(array []int) *Node {
	var bt BinaryTree
	root := bt.Insert(nil, "left", array[0])
	for i := 1; i < len(array); i++ {
		side, node := find(array[i], root)
		bt.Insert(node, side, array[i])
	}
	return root
}

func find(value int, root *Node) (side string, node *Node) {
	next := root
	for next != nil {
		node = next
		if value <= next.value {
			side = "left"
			next = next.left
		} else {
			side = "right"
			next = next.right
		}
	}
	return
}

func Heapsort(array *IntHeap) []int {
	heap.Init(array)
	n := array.Len()
	sortedArray := make([]int, n)
	for i := n - 1; array.Len() > 0; i-- {
		sortedArray[i] = heap.Pop(array).(int)
	}
	return sortedArray
}
