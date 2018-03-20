package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
)

func main() {
	graph := graphs.ReadGraph()
	dfs.DFS(graph)
}
