package main

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/graphs/bfs"
)

func main() {
	var s string
	graph := graphs.ReadGraph()
	fmt.Println("Enter the initial node:")
	fmt.Scanf("%s", &s)
	bfs.BFS(graph, s)
}
