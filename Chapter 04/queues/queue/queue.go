package queue

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

func (q *Queue) Enqueue(v int) {
	q.queue.PushFront(v)
}

func (q *Queue) Dequeue() int {
	if q.queue.Len() == 0 {
		return -1
	}
	element := q.queue.Back()
	q.queue.Remove(element)
	return element.Value.(int)
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
