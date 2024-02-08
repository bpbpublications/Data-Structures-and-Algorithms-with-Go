package graph

import (
	"container/heap"
)

type WeightedEdge struct {
	u, v   Node
	weight int
}

type WeightedGraph struct {
	nodes map[Node]struct{}
	edges map[WeightedEdge]struct{}
}

func NewWeightedGraph() WeightedGraph {
	return WeightedGraph{
		nodes: make(map[Node]struct{}),
		edges: make(map[WeightedEdge]struct{}),
	}
}

func (wg *WeightedGraph) AddNode(n Node) {
	wg.nodes[n] = struct{}{}
}

func (wg *WeightedGraph) AddEdgee(u, v Node, w int) {
	e := WeightedEdge{u, v, w}
	wg.edges[e] = struct{}{}
}

func (wg *WeightedGraph) RemoveNode(n Node) {
	delete(wg.nodes, n)
	for e := range wg.edges {
		if e.u == n || e.v == n {
			delete(wg.edges, e)
		}
	}
}

func Prim(wg *WeightedGraph, start *Node) *MST {
	treeNodes := make(map[Node]struct{})
	treeEdges := make(map[WeightedEdge]struct{})
	treeNodes[*start] = struct{}{}
	for len(treeNodes) != len(wg.nodes) {
		node, minEdge := minEdge(wg, treeNodes)
		treeNodes[node] = struct{}{}
		treeEdges[minEdge] = struct{}{}
	}
	return &MST{treeNodes, treeEdges}
}

func minEdge(wg *WeightedGraph, nodes map[Node]struct{}) (n Node, we WeightedEdge) {
	min := 1000
	for e := range wg.edges {
		_, ok1 := nodes[e.u]
		_, ok2 := nodes[e.v]
		if ok1 && !ok2 && e.v.value < min {
			n = e.v
			we = e
		}
	}
	return
}

func Kruskal(wg *WeightedGraph) *MST {
	treeEdges := make(map[WeightedEdge]struct{})
	pq := make(PriorityQueue, 0)
	for edge := range wg.edges {
		heap.Push(&pq, edge)
	}

	// create forest
	forest := make(map[int][]Node)
	i := 0
	for node := range wg.nodes {
		forest[i] = append(forest[i], node)
		i++
	}

	for n := 0; n < len(wg.nodes)-1 || pq.Len() == 0; {
		edge := heap.Pop(&pq).(Element).value
		i := findInForest(forest, edge.u.value)
		j := findInForest(forest, edge.v.value)
		if i != j {
			treeEdges[edge] = struct{}{}
			forest[i] = append(forest[i], forest[j]...)
			delete(forest, j)
			n++
		}
	}
	return &MST{wg.nodes, treeEdges}
}

func findInForest(forest map[int][]Node, node int) int {
	for i, v := range forest {
		// check all nodes in tree
		for _, n := range v {
			if n.value == node {
				return i
			}
		}
	}
	return -1
}

func TopSortWG(graph WeightedGraph) []Node {
	nodes := make(map[Node]struct{})
	edges := make(map[WeightedEdge]struct{})
	for k, v := range graph.nodes {
		nodes[k] = v
	}
	for k, v := range graph.edges {
		edges[k] = v
	}
	n := len(nodes)
	t := make([]Node, n)
	for i := 0; i < n; i++ {
		zeroIndegreeNode := findZeroIndegreeNodeWG(nodes, edges)
		t[i] = zeroIndegreeNode
		delete(nodes, zeroIndegreeNode)
		for edge := range edges {
			if edge.u.value == zeroIndegreeNode.value {
				delete(edges, edge)
			}
		}
	}
	return t
}

func findZeroIndegreeNodeWG(nodes map[Node]struct{}, edges map[WeightedEdge]struct{}) Node {
	for node := range nodes {
		isZeroDegree := true
		for edge := range edges {
			if edge.v.value == node.value {
				isZeroDegree = false
			}
		}
		if isZeroDegree {
			return node
		}
	}
	return Node{}
}

func CriticalPath(graph WeightedGraph) ([]int, []int, []int) {
	t := TopSortWG(graph)
	n := len(t)
	est := make([]int, n)
	est[0] = 0
	for i := 1; i < n; i++ {
		k := t[i].value
		est[k] = findMax(k, est, graph.edges)
	}
	lst := make([]int, n)
	lst[n-1] = est[n-1]
	for i := n - 2; i >= 0; i-- {
		k := t[i].value
		lst[k] = findMinLST(k, lst, graph.edges)
	}

	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = lst[i] - est[i]
	}

	return est, lst, l
}

func findMax(nodeValue int, est []int, edges map[WeightedEdge]struct{}) int {
	max := -1
	// find predecessors
	predecessors := make(map[int]int)
	for edge := range edges {
		if edge.v.value == nodeValue {
			predecessors[edge.u.value] = edge.weight
		}
	}

	// find maximum
	for p, v := range predecessors {
		if est[p]+v > max {
			max = est[p] + v
		}
	}

	return max
}

func findMinLST(nodeValue int, lst []int, edges map[WeightedEdge]struct{}) int {
	min := INF
	// find successors
	successors := make(map[int]int)
	for edge := range edges {
		if edge.u.value == nodeValue {
			successors[edge.v.value] = edge.weight
		}
	}

	// find minimum
	for p, v := range successors {
		if lst[p]-v < min {
			min = lst[p] - v
		}
	}

	return min
}
