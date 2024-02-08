package priorityqueue

type Element struct {
	value int
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value < pq[j].value
}

func (pq PriorityQueue) Swap(i, j int) {
	if pq.Len() == 0 {
		return
	}
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(v any) {
	element := Element{
		value: v.(int),
	}

	*pq = append(*pq, element)
}

func (pq *PriorityQueue) Pop() any {
	if pq.Len() == 0 {
		return -1
	}

	queue := *pq
	n := pq.Len() - 1
	element := queue[n]
	*pq = queue[0:n]
	return element
}
