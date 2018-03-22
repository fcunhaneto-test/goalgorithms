package dfs

import (
	"fmt"
	"goalgorithms/graphs"
	"goalgorithms/linklist"
)

func printDFS(tail linklist.LL) {
	n := tail.N.(*graphs.Vertex)
	fmt.Println("*********************************")
	fmt.Println("Vertex:", n.V)
	fmt.Println("CC:", n.CC)
	if n.P != nil {
		fmt.Println("Predecessor:", n.P.V)
	} else {
		fmt.Println("Predecessor: null")
	}
	fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
	fmt.Println("*********************************")
	for tail.Next != nil {
		tail = tail.Next
		if tail.Next != nil {
			n = tail.N.(*graphs.Vertex)
			fmt.Println("Vertex:", n.V)
			fmt.Println("CC:", n.CC)
			if n.P != nil {
				fmt.Println("Predecessor:", n.P.V)
			} else {
				fmt.Println("Predecessor: null")
			}
			fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
			fmt.Println("*********************************")
		}
	}
}
