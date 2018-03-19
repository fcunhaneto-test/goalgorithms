package tlogicalsort

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/llQueue"
)

var gtime int

// TopSort with Depth First Search (Busca em Profundidade)
func TopSort(graph map[string]*graphs.Node) {
	for _, g := range graph {
		if g.C == 0 {
			visit(graph, g)
		}
	}

	v := llQueue.LlPop()
	for v != nil {
		n := v.(*graphs.Node)
		fmt.Println(n.N)
		v = llQueue.LlPop()
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
	llQueue.LlPush(u)
}
