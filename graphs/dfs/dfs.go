package dfs

import (
	"goalgorithms/graphs"
	"goalgorithms/linklist"

	"terminal"
)

var cc int
var ga linklist.LL

// DFS Depth First Search (Busca em Profundidade)
func DFS(graph linklist.LL) linklist.LL {
	ga = ga.LlStart()
	for graph.Next != nil {
		g := graph.N.(*graphs.Vertex)
		if g.C == 0 {
			cc++
			g.CC = cc
			explore(g)
		}
		graph = graph.Next
	}

	terminal.Clear()
	return ga.GetHead()
}
