package main

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/graphs/bfs"
)

func main() {
	var s string
	graph := graphs.ReadGraph()
	fmt.Println("Enter the start vertex:")
	fmt.Scanf("%s", &s)

	bfs.BFS(s)

	for graph.Next != nil {
		n := graph.N.(*graphs.Vertex)
		fmt.Println("Vertex:", n.V)
		if n.P != nil {
			fmt.Println("Predecessor:", n.P.V)
		} else {
			fmt.Println("Predecessor: null")
		}
		fmt.Printf("Distance: %d\n", n.TI)
		fmt.Println("*********************************")

		graph = graph.Next
	}
}
