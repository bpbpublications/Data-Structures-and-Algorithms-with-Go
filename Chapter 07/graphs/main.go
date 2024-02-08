package main

import (
	"fmt"
	"graphs/graph"
)

func main() {
	node27 := graph.NewNode(27)
	node18 := graph.NewNode(18)
	node21 := graph.NewNode(21)
	node9 := graph.NewNode(9)
	node5 := graph.NewNode(5)
	node25 := graph.NewNode(25)

	g := graph.New()
	g.AddNode(node27)
	g.AddNode(node18)
	g.AddNode(node21)
	g.AddNode(node9)
	g.AddNode(node5)
	g.AddNode(node25)
	g.AddEdge(node27, node18)
	g.AddEdge(node27, node21)
	g.AddEdge(node27, node9)
	g.AddEdge(node18, node5)
	g.AddEdge(node21, node5)
	g.AddEdge(node21, node25)
	g.AddEdge(node9, node25)
	g.AddEdge(node5, node25)

	fmt.Println("BFS")
	graph.BFS(g, &node27)
	fmt.Println("DFS")
	graph.DFS(g, &node27)

	wg := graph.NewWeightedGraph()
	wg.AddNode(node21)
	wg.AddNode(node27)
	wg.AddNode(node18)
	wg.AddNode(node5)
	wg.AddNode(node9)
	wg.AddEdgee(node21, node27, 5)
	wg.AddEdgee(node21, node18, 10)
	wg.AddEdgee(node27, node18, 25)
	wg.AddEdgee(node27, node5, 15)
	wg.AddEdgee(node18, node9, 20)
	wg.AddEdgee(node5, node9, 30)

	mst := graph.Prim(&wg, &node21)
	mst.Print()

	mst = graph.Kruskal(&wg)
	mst.Print()

	a := [][]bool{
		{false, true, false, true, false},
		{false, false, false, false, true},
		{false, false, false, true, true},
		{false, false, false, false, false},
		{false, false, false, true, false},
	}
	p := graph.Warshall(a)
	fmt.Println(p)

	const INF = 99999

	w := [][]int{
		{0, 5, INF, 7, INF},
		{INF, 0, INF, INF, 10},
		{INF, INF, 0, 3, 6},
		{INF, INF, INF, 0, INF},
		{INF, INF, INF, 9, 0},
	}
	d := graph.Floyd(w)
	fmt.Println(d)

	w = [][]int{
		{0, 5, INF, 7, INF},
		{INF, 0, INF, INF, 10},
		{INF, INF, 0, INF, INF},
		{INF, INF, 3, 0, INF},
		{INF, INF, 6, 9, 0},
	}
	dd := graph.Dijkstra(0, w)
	fmt.Println(dd)

	node0 := graph.NewNode(0)
	node1 := graph.NewNode(1)
	node2 := graph.NewNode(2)
	node3 := graph.NewNode(3)
	node4 := graph.NewNode(4)

	fg := graph.NewFlowGraph()
	fg.AddNode(node0)
	fg.AddNode(node1)
	fg.AddNode(node2)
	fg.AddNode(node3)
	fg.AddNode(node4)
	fg.AddEdge(node0, node1, 5)
	fg.AddEdge(node0, node3, 7)
	fg.AddEdge(node1, node4, 10)
	fg.AddEdge(node3, node2, 3)
	fg.AddEdge(node4, node2, 6)
	fg.AddEdge(node4, node3, 9)

	maxFlow := graph.FordFulkerson(fg, &node0, &node2)
	fmt.Println(maxFlow)

	tsGraph := graph.New()
	tsGraph.AddNode(node0)
	tsGraph.AddNode(node1)
	tsGraph.AddNode(node2)
	tsGraph.AddNode(node3)
	tsGraph.AddNode(node4)
	tsGraph.AddEdge(node0, node1)
	tsGraph.AddEdge(node0, node3)
	tsGraph.AddEdge(node1, node4)
	tsGraph.AddEdge(node3, node2)
	tsGraph.AddEdge(node4, node2)
	tsGraph.AddEdge(node4, node3)
	t := graph.TopSort(tsGraph)
	fmt.Println(t)

	wg = graph.NewWeightedGraph()
	wg.AddNode(node0)
	wg.AddNode(node1)
	wg.AddNode(node2)
	wg.AddNode(node3)
	wg.AddNode(node4)
	wg.AddNode(node5)
	wg.AddEdgee(node0, node1, 5)
	wg.AddEdgee(node0, node2, 7)
	wg.AddEdgee(node1, node3, 2)
	wg.AddEdgee(node1, node4, 6)
	wg.AddEdgee(node2, node3, 4)
	wg.AddEdgee(node2, node4, 9)
	wg.AddEdgee(node3, node4, 3)
	wg.AddEdgee(node3, node5, 8)
	wg.AddEdgee(node4, node5, 6)
	est, lst, l := graph.CriticalPath(wg)
	fmt.Println(est, lst, l)
}
