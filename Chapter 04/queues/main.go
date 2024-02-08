package main

import (
	"container/heap"
	"fmt"
	"queues/priorityqueue"
	"queues/queue"
)

func main() {
	q := queue.New()

	q.Enqueue(27)
	q.Enqueue(5)
	q.Enqueue(1)
	q.Enqueue(18)
	q.Print()

	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	fmt.Println(q.Dequeue())
	q.Print()

	// Priority Queue
	pq := make(priorityqueue.PriorityQueue, 0)

	heap.Push(&pq, 27)
	heap.Push(&pq, 5)
	heap.Push(&pq, 1)
	heap.Push(&pq, 18)
	fmt.Println(pq)

	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(heap.Pop(&pq))
	fmt.Println(pq)
}
