package dfs

import (
	"goalgorithms/graphs"
	"goalgorithms/llQueue"
)

var gtime int

var tail llQueue.LL

// explore given the initial node exploits the descents of this node
func explore(graph map[string]*graphs.Node, u *graphs.Node) {
	gtime++
	u.TI = gtime
	u.C = 1
	u.CC = cc

	for i := 0; i < len(u.Adj); i++ {
		if u.Adj[i] == nil {
			continue
		}
		if u.Adj[i].C == 0 {
			u.Adj[i].P = u
			u.Adj[i].CC = cc
			u.C = 1
			explore(graph, u.Adj[i])
		}
	}

	u.C = 2
	gtime++
	u.TF = gtime
	tail = llQueue.LlPush(u)
}
