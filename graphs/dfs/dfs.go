package dfs

import (
	"goalgorithms/graphs"
	"goalgorithms/linklist"

	"terminal"
)

var cc int

// DFS Depth First Search (Busca em Profundidade)
func DFS(graph linklist.LL) {
	current := graph
	for current.Next != nil {
		g := current.N.(*graphs.Vertex)
		if g.C == 0 {
			cc++
			g.CC = cc
			explore(g)
		}
		current = current.Next
	}

	terminal.Clear()
	printDFS(linklist.GetHead())
}
