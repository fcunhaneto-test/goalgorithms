package main

import (
	"fmt"
	"goalgorithms/graphs"
)

func main() {
	graph := graphs.ReadGraph()

	for _, n := range graph {
		fmt.Println(*n)
	}
}
