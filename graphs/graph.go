package graphs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Node BFS of graph who values are integer
// type Node struct {
// 	N   string
// 	Adj []*Node
// 	C   int
// 	D   int
// 	P   *Node
// }

// Node DFS of graph who values are integer
type Node struct {
	N   string
	Adj []*Node
	C   int
	TI  int
	TF  int
	P   *Node
}

// Graph an array of all graphnodes
var Graph map[string]*Node

/*
ReadGraph first read an integer representing a number of nodes in graph, then read an string representing node, then reads strings on the same line separated by space representing the adjacent nodes.
Example:
Enter the number of nodes:
5
Enter node:
A
Enter the adjacent nodes:
B C F
*/
func ReadGraph() map[string]*Node {
	var num int
	var s string
	var line []string

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the number of nodes: ")
	fmt.Scanf("%d", &num)

	Graph = make(map[string]*Node)
	a := make(map[string][]string)

	for i := 0; i < num; i++ {
		fmt.Println("Enter node: ")
		fmt.Scanf("%s", &s)
		node := new(Node)
		node.N = s
		Graph[s] = node

		fmt.Printf("Enter the adjacent nodes for node: %s\n", s)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]

		if text != "" {
			line = strings.Split(string(text), " ")
			a[s] = line
			// fmt.Println("s:", s, "a[s]:", a[s])
		} else {
			a[s] = nil
		}
	}

	for c, n := range Graph {
		if a[c] != nil {
			b := a[c]
			n.Adj = make([]*Node, len(b))
			for i := 0; i < len(b); i++ {
				n.Adj[i] = Graph[b[i]]
			}
		}
	}

	return Graph
}
