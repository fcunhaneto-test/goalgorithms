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

// Gmap map of graph
var Gmap map[string]*Vertex

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
	var graph linklist.LL
	var line []string

	reader := bufio.NewReader(os.Stdin)

	graph = graph.LlStart()
	Gmap = make(map[string]*Vertex)

	fmt.Println("Enter the vertex name (or 0 to stop):")
	fmt.Scanf("%s", &s)
	for strings.Compare(s, "0") != 0 {
		v = new(Vertex)
		v.V = s
		graph = graph.LlEnqueue(v)
		Gmap[s] = v
		fmt.Println("Enter the vertex name (or nothing to stop):")
		fmt.Scanf("%s", &s)
	}

	graph = graph.GetHead()
	for graph.Next != nil {
		v = graph.N.(*Vertex)
		fmt.Printf("Enter the adjacent nodes for node: %s\n", v.V)
		text, _ := reader.ReadString('\n')
		text = text[:len(text)-1]
		line = strings.Split(string(text), " ")

		v.Adj = make([]*Vertex, len(line))
		for i := 0; i < len(line); i++ {
			v.Adj[i] = Gmap[line[i]]
		}

		graph = graph.Next

	}

	graph = graph.GetHead()

	return graph
}
