package dfs

import (
	"goalgorithms/graphs"

	"terminal"
)

var cc int

// DFS Depth First Search (Busca em Profundidade)
func DFS(graph map[string]*graphs.Node, s string) {
	cc++
	graph[s].CC = cc
	explore(graph, graph[s])
	for _, g := range graph {
		if g.C == 0 {
			cc++
			graph[s].CC = cc
			explore(graph, g)
		}
	}

	terminal.Clear()
	printDFS()
}
