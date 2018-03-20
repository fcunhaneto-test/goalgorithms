package dfs

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/llQueue"
)

var gtime int

var tail llQueue.LL

func explore(graph map[string]*graphs.Node, u *graphs.Node) {
	gtime++
	u.TI = gtime
	u.C = 1

	for i := 0; i < len(u.Adj); i++ {
		if u.Adj[i] == nil {
			continue
		}
		if u.Adj[i].C == 0 {
			u.Adj[i].P = u
			explore(graph, u.Adj[i])
		}
	}

	u.C = 2
	gtime++
	u.TF = gtime
	tail = llQueue.LlPush(u)
}

func printDFS() {
	var n *graphs.Node
	n = tail.V.(*graphs.Node)
	fmt.Println("*********************************")
	fmt.Println("Node:", n.N)
	if n.P != nil {
		fmt.Println("Predecessor:", n.P.N)
	} else {
		fmt.Println("Predecessor: null")
	}
	fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
	fmt.Println("*********************************")
	for tail.Previ != nil {
		tail = tail.Previ
		if tail.Previ != nil {
			n = tail.V.(*graphs.Node)
			fmt.Println("Node:", n.N)
			if n.P != nil {
				fmt.Println("Predecessor:", n.P.N)
			} else {
				fmt.Println("Predecessor: null")
			}
			fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
			fmt.Println("*********************************")
		}
	}
}
