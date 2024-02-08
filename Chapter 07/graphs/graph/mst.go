package graph

import "fmt"

type MST struct {
	nodes map[Node]struct{}
	edges map[WeightedEdge]struct{}
}

func (m MST) Print() {
	fmt.Println("Nodes:")
	for n := range m.nodes {
		fmt.Printf("%v ", n.value)
	}
	fmt.Println()
	fmt.Println("Edges: ")
	for e := range m.edges {
		fmt.Printf("%v ", e)
	}
}
