package main

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
)

func main() {
	var s string
	graph := graphs.ReadGraph()
	fmt.Println("Enter the initial node:")
	fmt.Scanf("%s", &s)
	dfs.DFS(graph, s)
}
