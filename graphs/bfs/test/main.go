package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/bfs"
)

func main() {
	var graph graphs.Graph
	graph = graphs.ReadGraph()
	bfs.BFS(graph, graph.G[0])
}
