package dfs

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/llQueue"
	"terminal"
)

// TopSort with Depth First Search (Busca em Profundidade)
func TopSort(graph map[string]*graphs.Node) {
	for _, g := range graph {
		if g.C == 0 {
			explore(graph, g)
		}
	}

	terminal.Clear()
	printDFS()
	v := llQueue.LlPop()
	for v != nil {
		n := v.(*graphs.Node)
		fmt.Println(n.N)
		v = llQueue.LlPop()
	}
}
