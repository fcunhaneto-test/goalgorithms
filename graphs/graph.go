package graphs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Node of graph who values are integer
type Node struct {
	N   string
	Adj []*Node
	C   int
	D   int
	P   *Node
}

// Graph an array of all graphnodes
var Graph map[string]*Node

func readNode(n *Node) {
	var node Node
	var line []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Enter the adjacent nodes for node: %s\n", n.N)
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	if text != "" {
		line = strings.Split(string(text), " ")
		n.Adj = make([]*Node, len(line))
		for i := 0; i < len(line); i++ {
			n.Adj[i] = Graph[line[i]]
		}
	} else {
		node.Adj = nil
	}
}

/*
ReadGraph first read an integer representing a number of nodes in graph, then read integer representing node number, then reads integers on the same space-separated line representing the adjacent nodes.
Example:
# Enter the number of nodes:
# 5
#"Enter node:
# 1
# Enter the adjacent nodes:
# 3 4 5
*/
func ReadGraph() map[string]*Node {
	var num int
	var s string
	fmt.Println("Enter the number of nodes: ")
	fmt.Scanf("%d", &num)

	Graph = make(map[string]*Node)

	for i := 0; i < num; i++ {
		fmt.Println("Enter node: ")
		fmt.Scanf("%s", &s)
		node := new(Node)
		node.N = s
		Graph[s] = node
	}

	for _, n := range Graph {
		readNode(n)
	}

	return Graph
}

/*
lineToInteger transforms an array of strings into an array of float64
param:
line: []string
return:
s: []float64
*/
func lineToInteger(line []string) []int {
	var num int
	var err error
	var s = []int{}

	for _, l := range line {
		num, err = strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		s = append(s, num)
	}

	return s
}
