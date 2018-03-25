package tsort

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
	"goalgorithms/linklist"
)

// Tsort Topological Sort
func Tsort(graph linklist.LL) {
	g := dfs.DFS(graph)
	for g.Next != nil {
		v := g.N.(*graphs.Vertex)
		fmt.Println(v.V)
		g = g.Next
	}
}
