package graph

type Element struct {
	value WeightedEdge
}

type PriorityQueue []Element

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].value.weight < pq[j].value.weight
}

func (pq PriorityQueue) Swap(i, j int) {
	if pq.Len() == 0 {
		return
	}
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(v any) {
	element := Element{
		value: v.(WeightedEdge),
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
