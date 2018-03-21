package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
)

func main() {
	graph := graphs.ReadGraph()
	dfs.TopSort(graph)

	// for _, n := range graph {
	// 	fmt.Println("Node:", n.N)
	// 	fmt.Println("Color:", n.C)
	// 	fmt.Printf("Time: %d/%d\n", n.TI, n.TF)
	// 	if n.P != nil {
	// 		fmt.Println("Predecessor:", n.P.N)
	// 	} else {
	// 		fmt.Println("Predecessor: null")
	// 	}
	// 	fmt.Println("*********************************")

	// }
}
