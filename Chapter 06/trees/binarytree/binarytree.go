package binarytree

import (
	"container/heap"
	"fmt"
)

type Node struct {
	value  int
	left   *Node
	right  *Node
	parent *Node
}

func (n *Node) DeleteLeaf() {
	if n.parent.left == n {
		n.parent.left = nil
	} else {
		n.parent.right = nil
	}
}

func (n *Node) findLeftmost() *Node {
	node := n
	next := n
	for next != nil {
		node = next
		if next.left != nil {
			next = next.left
		} else {
			next = next.right
		}
	}

	return node
}

type BinaryTree struct {
	root *Node
}

func (bt *BinaryTree) GetRoot() *Node {
	return bt.root
}

func (bt *BinaryTree) Insert(node *Node, side string, value int) *Node {
	newNode := &Node{value, nil, nil, nil}

	// inserting to empty tree
	if bt.root == nil {
		bt.root = newNode
		return bt.root
	}

	if side == "left" {
		if node.left == nil {
			// inserting a leaf
			node.left = newNode
		} else {
			// inserting internal node
			newNode.left = node.left
			node.left = newNode
		}
	} else {
		if node.right == nil {
			// inserting a leaf
			node.right = newNode
		} else {
			// inserting internal node
			newNode.right = node.right
			node.right = newNode
		}
	}

	newNode.parent = node
	return newNode
}

func (bt *BinaryTree) Delete(node *Node) {
	// delete leaf
	if node.left == nil && node.right == nil {
		// delete root when the root is the only node
		if node == bt.root {
			bt.root = nil
		} else {
			node.DeleteLeaf()
		}
		// delete internal node
	} else {
		leftmostNode := node.findLeftmost()
		node.value = leftmostNode.value
		leftmostNode.DeleteLeaf()
	}
}

func Preorder(node *Node) {
	if node != nil {
		fmt.Println(node.value)
		Preorder(node.left)
		Preorder(node.right)
	}
}

func Inorder(node *Node) {
	if node != nil {
		Inorder(node.left)
		fmt.Println(node.value)
		Inorder(node.right)
	}
}

func Postorder(node *Node) {
	if node != nil {
		Postorder(node.left)
		Postorder(node.right)
		fmt.Println(node.value)
	}
}

func Levelorder(node *Node) {
	next := node
	queue := New()
	queue.Enqueue(node)
	for !queue.IsEmpty() {
		next = queue.Dequeue()
		fmt.Println(next.value)
		if next.left != nil {
			queue.Enqueue(next.left)
		}
		if next.right != nil {
			queue.Enqueue(next.right)
		}
	}
}

type WeightNode struct {
	value  int
	weight int
	left   *WeightNode
	right  *WeightNode
}

type WeightBinnaryTree struct {
	root *WeightNode
}

func Huffman(w []int, e int) *WeightBinnaryTree {
	pq := make(PriorityQueue, 0)
	for i := 0; i < e; i++ {
		node := WeightNode{
			value:  i,
			weight: w[i],
			left:   nil,
			right:  nil,
		}
		heap.Push(&pq, node)
	}

	var nodeZ WeightNode
	for i := 0; i < e-1; i++ {
		nodeX := heap.Pop(&pq).(Element).value
		nodeY := heap.Pop(&pq).(Element).value
		nodeZ = WeightNode{
			value:  e,
			weight: nodeX.weight + nodeY.weight,
			left:   &nodeX,
			right:  &nodeY,
		}
		heap.Push(&pq, nodeZ)
	}

	return &WeightBinnaryTree{
		root: &nodeZ,
	}
}
