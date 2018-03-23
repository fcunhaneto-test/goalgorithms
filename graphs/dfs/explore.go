package dfs

import (
	"goalgorithms/graphs"
	"goalgorithms/linklist"
)

var gtime int

var tail linklist.LL

// explore given the initial node exploits the descents of this node
func explore(u *graphs.Vertex) {
	gtime++
	u.TI = gtime
	u.C = 1

	for i := 0; i < len(u.Adj); i++ {
		if u.Adj[i] == nil {
			continue
		}
		if u.Adj[i].C == 0 {
			u.Adj[i].P = u
			u.Adj[i].CC = cc
			u.C = 1
			explore(u.Adj[i])
		}
	}

	u.C = 2
	gtime++
	u.TF = gtime

	ga = ga.LlPush(u)
}
