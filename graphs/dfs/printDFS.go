package dfs

import (
	"fmt"
	"goalgorithms/graphs"
)

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
			fmt.Println("CC:", n.CC)
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
