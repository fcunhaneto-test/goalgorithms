package dfsfix

import (
	"goalgorithms/graphs"
)

var gtime int

// DFS Depth First Search (Busca em Profundidade)
func DFS(graph map[string]*graphs.Node, s string) {
	visit(graph, graph[s])
	for _, g := range graph {
		if g.C == 0 {
			visit(graph, g)
		}
	}
}

func visit(graph map[string]*graphs.Node, u *graphs.Node) {
	gtime++
	u.TI = gtime
	u.C = 1

	for i := 0; i < len(u.Adj); i++ {
		if u.Adj[i] == nil {
			continue
		}
		if u.Adj[i].C == 0 {
			u.Adj[i].P = u
			visit(graph, u.Adj[i])
		}
	}

	u.C = 2
	gtime++
	u.TF = gtime
}
