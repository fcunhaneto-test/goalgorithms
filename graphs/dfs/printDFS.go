package dfs

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/linklist"
)

func PrintDFS(head linklist.LL) {
	var n *graphs.Vertex
	for head.Next != nil {
		n = head.N.(*graphs.Vertex)
		fmt.Println("Vertex:", n.V)
		fmt.Println("CC:", n.CC)
		if n.P != nil {
			fmt.Println("Predecessor:", n.P.V)
		} else {
			fmt.Println("Predecessor: null")
		}
		fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
		fmt.Println("*********************************")

		head = head.Next
	}
}
