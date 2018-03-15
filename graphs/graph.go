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
	N   int
	Adj []int
}

type Graph struct {
	G []Node
}

func readNode(num int) Node {
	var a []int
	var node Node
	var line []string
	reader := bufio.NewReader(os.Stdin)

	node.N = num
	fmt.Println("Enter the adjacent nodes: ")
	text, _ := reader.ReadString('\n')
	text = text[:len(text)-1]

	if text != "" {
		line = strings.Split(string(text), " ")
		a = lineToInteger(line)
		node.Adj = a
	} else {
		node.Adj = nil
	}

	return node
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
func ReadGraph() Graph {
	var n, num int
	var graph Graph
	fmt.Println("Enter the number of nodes: ")
	fmt.Scanf("%d", &num)

	graph.G = make([]Node, num)

	for i := 0; i < num; i++ {
		fmt.Println("Enter node: ")
		fmt.Scanf("%d", &n)
		node := readNode(n)
		graph.G[i] = node
	}

	return graph
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
