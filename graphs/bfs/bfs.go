package bfs

import (
	"goalgorithms/graphs"
	"goalgorithms/linklist"
)

// BFS Breadth First Traversal (Busca em Largura)
func BFS(s string) {
	var ll linklist.LL

	v := graphs.Gmap[s]

	v.C = 1

	ll = ll.LlStart()
	ll = ll.LlEnqueue(v)

	for !ll.LlEmpty() {
		u := ll.LlPop()
		w := u.(*graphs.Vertex)
		t := w.Adj

		if t != nil {
			for i := 0; i < len(t); i++ {
				if t[i].C == 0 {
					t[i].C = 1
					t[i].TI = w.TI + 1
					t[i].P = w
					ll = ll.LlEnqueue(t[i])
				}
			}
		}
	}

}
