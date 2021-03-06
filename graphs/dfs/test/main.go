package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
)

func main() {
	graph := graphs.ReadGraph()
	graph = dfs.DFS(graph)
	dfs.PrintDFS(graph)
}
