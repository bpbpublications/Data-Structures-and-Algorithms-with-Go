package graph

type FlowEdge struct {
	u, v     Node
	capacity int
	flow     int
}

type FlowGraph struct {
	nodes map[Node]struct{}
	edges map[FlowEdge]struct{}
}

func NewFlowGraph() *FlowGraph {
	return &FlowGraph{
		nodes: make(map[Node]struct{}),
		edges: make(map[FlowEdge]struct{}),
	}
}

func (fg *FlowGraph) AddNode(n Node) {
	fg.nodes[n] = struct{}{}
}

func (fg *FlowGraph) AddEdge(u, v Node, c int) {
	e := FlowEdge{u, v, c, 0}
	fg.edges[e] = struct{}{}
}

func (fg *FlowGraph) RemoveNode(n Node) {
	delete(fg.nodes, n)
	for e := range fg.edges {
		if e.u == n || e.v == n {
			delete(fg.edges, e)
		}
	}
}

func AugmentingPath(g *FlowGraph, source *Node, target *Node) (ok bool, path []FlowEdge) {
	ok = false
	visit := make(map[int]bool)
	for n := range g.nodes {
		visit[n.value] = false
	}
	visit[source.value] = true
	queue := NewQueue()
	queue.Enqueue(source)
	for !queue.IsEmpty() {
		u := queue.Dequeue()
		for edge := range g.edges {
			if edge.u.value == u.value && !visit[edge.v.value] && edge.capacity != edge.flow {
				if len(path) == 0 || !inPath(edge.u.value, path) && inPathDest(edge.u.value, path) {
					visit[edge.v.value] = true
					path = append(path, edge)
					n := edge.v
					queue.Enqueue(&n)
					if edge.v.value == target.value {
						ok = true
						return
					}
				}
			}
		}
	}
	return
}

func inPath(node int, path []FlowEdge) bool {
	for _, edge := range path {
		if edge.u.value == node {
			return true
		}
	}
	return false
}

func inPathDest(node int, path []FlowEdge) bool {
	for _, edge := range path {
		if edge.v.value == node {
			return true
		}
	}
	return false
}

func FordFulkerson(graph *FlowGraph, source, target *Node) (maxFLow int) {
	for e := range graph.edges {
		e.flow = 0
	}

	ok, path := AugmentingPath(graph, source, target)
	for ok {
		cf := pathFLow(path)
		for _, edge := range path {
			delete(graph.edges, edge)
			edge.flow += cf
			graph.edges[edge] = struct{}{}
		}
		maxFLow += cf
		ok, path = AugmentingPath(graph, source, target)
	}
	return
}

func pathFLow(path []FlowEdge) (flow int) {
	flow = INF
	for _, edge := range path {
		cf := edge.capacity - edge.flow
		if cf < flow {
			flow = cf
		}
	}
	return
}
