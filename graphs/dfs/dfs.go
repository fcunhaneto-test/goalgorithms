package dfs

import (
	"goalgorithms/graphs"
	"terminal"
)

// DFS Depth First Search (Busca em Profundidade)
func DFS(graph map[string]*graphs.Node) {
	for _, g := range graph {
		if g.C == 0 {
			explore(graph, g)
		}
	}

	terminal.Clear()
	printDFS()
}
