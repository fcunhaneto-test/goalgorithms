package bfs

import (
	"goalgorithms/graphs"
	"goalgorithms/linklist"
	"goalgorithms/llqueue"
)

func BFS(graph linklist.LL, s string) {
	// var ll linklist.LL
	u := linklist.LlFindByS(s, 0)

	v := u.N.(*graphs.Vertex)
	v.C = 1

	ll := linklist.LlStart()
	ll = ll.LlPush(v)

	for ll.

}
