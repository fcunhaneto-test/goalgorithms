package bfs

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/llQueue"
)

// BFS Breadth First Search (Busca em Largura)
func BFS(graph map[string]*graphs.Node, name string) {
	node := graph[name]
	node.C = 1

	llQueue.LlPush(node)

	for !llQueue.LlEmpty() {
		aux := llQueue.LlPop()
		u := aux.(*graphs.Node)
		for i := 0; i < len(u.Adj); i++ {
			if u.Adj[i].C == 0 {
				u.Adj[i].C = 1
				u.Adj[i].D = u.D + 1
				u.Adj[i].P = u
				llQueue.LlPush(u.Adj[i])
			}
		}
		u.C = 2
	}

	for _, n := range graph {
		fmt.Println("Node:", n.N)
		fmt.Println("Color:", n.C)
		fmt.Println("Distance:", n.D)
		if n.P != nil {
			fmt.Println("Predecessor:", n.P.N)
		} else {
			fmt.Println("Predecessor: null")
		}
		fmt.Println("*********************************")

	}
}
