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
	n   int
	adj []int
}

// Nodes all nodes of graph
type Nodes []Node

/*
ReadNode first read an integer representing a node in graph then reads integers on the same space-separated line representing the adjacent nodes.
Example:
# Enter vertex:
# 1
# Enter the adjacent vertices:
# 2 6 7 10 5
*/
func ReadNode() Node {
	var num int
	var a []int
	var node Node
	var line []string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter vertex:")
	fmt.Scanf("%d", &num)
	node.n = num

	fmt.Println("Enter the adjacent vertices: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	line = strings.Split(string(text), " ")
	a = lineToInteger(line)
	node.adj = a

	return node
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

func BFS(node Nodes, n Node) {

}
