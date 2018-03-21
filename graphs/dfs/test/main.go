package main

import (
	"goalgorithms/graphs"
	"goalgorithms/graphs/dfs"
)

func main() {
	// var s string
	ll := graphs.ReadGraph()
	// fmt.Println("Enter the initial node:")
	// fmt.Scanf("%s", &s)
	dfs.DFS(ll)

}
