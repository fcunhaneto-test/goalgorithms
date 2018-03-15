package main

import (
	"fmt"
	"goalgorithms/graphs"
)

func main() {
	var graph []graphs.Node
	graph = graphs.ReadGraph()

	for i := 0; i < len(graph); i++ {
		fmt.Println(graph)
	}
}
