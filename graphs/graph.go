package graphs

import (
	"bufio"
	"fmt"
	"goalgorithms/linklist"
	"os"
	"strings"
)

// Vertex struct that represent vertex of graph
type Vertex struct {
	V   string
	Adj []*Vertex
	CC  int
	C   int
	TI  int
	TF  int
	P   *Vertex
}

// Graph an array of all graphnodes
var Graph map[string]*Vertex

/*
ReadGraph first read vertex name in graph, then reads strings on the same line separated by space representing the name of adjacent vertices.
Example:
Enter node:
A
Enter the adjacent nodes:
B C F
*/
func ReadGraph() linklist.LL {
	var s string
	var v *Vertex
	var current linklist.LL
	var line []string

	Graph = make(map[string]*Vertex)
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the vertex name (or 0 to stop):")
	fmt.Scanf("%s", &s)
	for strings.Compare(s, "0") != 0 {
		v = new(Vertex)
		v.V = s
		current = linklist.LlEnqueue(v, current)
		Graph[s] = v
		fmt.Println("Enter the vertex name (or nothing to stop):")
		fmt.Scanf("%s", &s)
	}

	current = linklist.GetHead()
	for current.Next != nil {
		v = current.N.(*Vertex)
		fmt.Printf("Enter the adjacent nodes for node: %s\n", v.V)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		line = strings.Split(string(text), " ")

		v.Adj = make([]*Vertex, len(line))
		for i := 0; i < len(line); i++ {
			v.Adj[i] = Graph[line[i]]
		}

		current = current.Next

	}

	current = linklist.GetHead()
	for current.Next != nil {
		v = current.N.(*Vertex)
		fmt.Println(v.V, v.Adj)
		current = current.Next
	}

	current = linklist.GetHead()

	return current
}
