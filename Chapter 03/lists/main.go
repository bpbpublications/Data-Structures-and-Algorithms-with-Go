package main

import (
	"container/list"
	"container/ring"
	"fmt"
	"lists/linkedlist"
)

func main() {
	l := list.New()

	head := l.PushFront(5)
	tail := l.PushBack(27)

	l.InsertBefore(1, tail)
	l.InsertAfter(18, head)

	fmt.Println(l.Len())
	fmt.Print("[")
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Next() != nil {
			fmt.Printf("%v, ", node.Value)
		} else {
			fmt.Print(node.Value)
		}
	}
	fmt.Println("]")

	l.Remove(tail)

	fmt.Println(l.Len())
	fmt.Print("[")
	for node := l.Back(); node != nil; node = node.Prev() {
		if node.Prev() != nil {
			fmt.Printf("%v, ", node.Value)
		} else {
			fmt.Print(node.Value)
		}
	}
	fmt.Println("]")

	// Ring
	r := ring.New(4)
	fmt.Println(r.Len())

	for i := 0; i < r.Len(); i++ {
		r.Value = i
		r = r.Next()
	}

	r.Do(func(p any) {
		fmt.Println(p.(int))
	})

	// List
	l1 := linkedlist.New()
	l1.Insert(1)
	l1.Insert(18)
	l1.Insert(9)
	l1.Insert(21)

	l1.Print()

	l1.Remove(9)
	l1.Print()

	l2 := linkedlist.New()
	l2.InsertOrdered(5)
	l2.InsertOrdered(27)
	l2.InsertOrdered(25)
	l2.InsertOrdered(8)

	l2.Print()

	n1 := l1.Find(18)
	fmt.Println(n1)

	n2 := l1.Find(27)
	fmt.Println(n2)

	l1.Concatenate(l2)
	l1.Print()
}
