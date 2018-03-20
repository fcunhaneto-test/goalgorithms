package dfs

import (
	"goalgorithms/graphs"

	"terminal"
)

// DFSFix Depth First Search (Busca em Profundidade)
func DFSFix(graph map[string]*graphs.Node, s string) {
	explore(graph, graph[s])
	for _, g := range graph {
		if g.C == 0 {
			explore(graph, g)
		}
	}

	terminal.Clear()
	printDFS()
}
