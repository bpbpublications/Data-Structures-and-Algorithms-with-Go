package main

import (
	"fmt"
	"tree/binarytree"
)

func main() {
	// Create a tree
	var bt binarytree.BinaryTree
	node18 := bt.Insert(nil, "left", 18)
	node8 := bt.Insert(node18, "left", 8)
	node25 := bt.Insert(node18, "rigth", 25)
	node5 := bt.Insert(node8, "left", 5)
	bt.Insert(node8, "right", 9)
	bt.Insert(node25, "left", 21)
	bt.Insert(node25, "right", 27)
	bt.Delete(node25)
	bt.Delete(node5)

	fmt.Println("Preorder")
	binarytree.Preorder(bt.GetRoot())
	fmt.Println("Inorder")
	binarytree.Inorder(bt.GetRoot())
	fmt.Println("Postorder")
	binarytree.Postorder(bt.GetRoot())
	fmt.Println("Levelorder")
	binarytree.Levelorder(bt.GetRoot())
	array := &binarytree.IntHeap{18, 27, 5, 21, 1, 9}
	fmt.Println(binarytree.Heapsort(array))

	w := []int{27, 25, 18, 9, 21}
	binarytree.Huffman(w, len(w))
}
