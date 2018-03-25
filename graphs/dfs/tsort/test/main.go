package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs/tsort"
)

func main() {
	graph := graphs.ReadGraph()
	tsort.Tsort(graph)
}
