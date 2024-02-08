package binarytree

import (
	"container/list"
	"fmt"
)

type Queue struct {
	queue *list.List
}

func New() *Queue {
	return &Queue{queue: list.New()}
}

func (q *Queue) Enqueue(node *Node) {
	q.queue.PushFront(node)
}

func (q *Queue) Dequeue() *Node {
	if q.queue.Len() == 0 {
		return nil
	}
	element := q.queue.Back()
	q.queue.Remove(element)
	return element.Value.(*Node)
}

func (q *Queue) IsEmpty() bool {
	return q.queue.Len() == 0
}

func (q *Queue) Print() {
	fmt.Print("[")
	for e := q.queue.Front(); e != nil; e = e.Next() {
		if e.Next() != nil {
			fmt.Printf("%v, ", e.Value)
		} else {
			fmt.Print(e.Value)
		}
	}
	fmt.Println("]")
}
