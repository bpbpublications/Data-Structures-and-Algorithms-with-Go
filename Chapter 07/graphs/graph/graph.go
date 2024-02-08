package graph

type Node struct {
	value int
}

func NewNode(value int) Node {
	return Node{value}
}

type Edge struct {
	u, v Node
}

type Graph struct {
	nodes map[Node]struct{}
	edges map[Edge]struct{}
}

func New() *Graph {
	return &Graph{
		nodes: make(map[Node]struct{}),
		edges: make(map[Edge]struct{}),
	}
}

func (g *Graph) AddNode(n Node) {
	g.nodes[n] = struct{}{}
}

func (g *Graph) AddEdge(u, v Node) {
	e := Edge{u, v}
	g.edges[e] = struct{}{}
}

func (g *Graph) RemoveNode(n Node) {
	delete(g.nodes, n)
	for e := range g.edges {
		if e.u == n || e.v == n {
			delete(g.edges, e)
		}
	}
}

func BFS(g *Graph, start *Node) {
	visit := make(map[int]bool)
	for n := range g.nodes {
		visit[n.value] = false
	}
	visit[start.value] = true
	queue := NewQueue()
	queue.Enqueue(start)
	for !queue.IsEmpty() {
		u := queue.Dequeue()
		for edge := range g.edges {
			if edge.u.value == u.value && !visit[edge.v.value] {
				visit[edge.v.value] = true
				n := edge.v
				queue.Enqueue(&n)
			}
		}
	}
}

func DFS(g *Graph, start *Node) {
	visit := make(map[int]bool)
	for n := range g.nodes {
		visit[n.value] = false
	}
	dfsVisit(g, start, visit)
}

func dfsVisit(g *Graph, u *Node, visit map[int]bool) {
	visit[u.value] = true
	for edge := range g.edges {
		if edge.u.value == u.value && !visit[edge.v.value] {
			dfsVisit(g, &edge.v, visit)
		}
	}
}

func Warshall(a [][]bool) (p [][]bool) {
	p = a
	for k := 0; k < len(p); k++ {
		for i := 0; i < len(p); i++ {
			for j := 0; j < len(p); j++ {
				p[i][j] = p[i][j] ||
					(p[i][k] && p[k][j])
			}
		}
	}
	return
}

func Floyd(w [][]int) (d [][]int) {
	d = w
	for k := 0; k < len(d); k++ {
		for i := 0; i < len(d); i++ {
			for j := 0; j < len(d); j++ {
				if d[i][j] > d[i][k]+d[k][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	return
}

const INF = 99999

func Dijkstra(start int, w [][]int) (d map[int]int) {
	n := len(w)

	s := make(map[int]struct{})
	s[start] = struct{}{}
	v := make(map[int]struct{})
	d = make(map[int]int)

	for i := 0; i < n; i++ {
		if i != start {
			v[i] = struct{}{}
			d[i] = w[start][i]
		}
	}

	for len(v) != 0 {
		i := findMin(d, v)
		s[i] = struct{}{}
		delete(v, i)
		for j := range v {
			if d[i]+w[i][j] < d[j] {
				d[j] = d[i] + w[i][j]
			}
		}
	}
	return
}

func findMin(d map[int]int, v map[int]struct{}) (node int) {
	min := INF
	for i := range d {
		_, ok := v[i]
		if d[i] <= min && ok {
			node = i
			min = d[i]
		}
	}
	return
}

func TopSort(graph *Graph) []Node {
	nodes := graph.nodes
	edges := graph.edges
	n := len(nodes)
	t := make([]Node, n)
	for i := 0; i < n; i++ {
		zeroIndegreeNode := findZeroIndegreeNode(nodes, edges)
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

func findZeroIndegreeNode(nodes map[Node]struct{}, edges map[Edge]struct{}) Node {
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
