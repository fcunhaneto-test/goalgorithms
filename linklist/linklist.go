package linklist

import "fmt"

// Node struct for link list
type Node struct {
	OB    interface{}
	Next  *Node
	Previ *Node
}

// LLInit init list link
func LLInit(v interface{}) *Node {
	node := &Node{
		OB:    v,
		Next:  nil,
		Previ: nil,
	}
	return node
}

// InsertLL insert struct in link list
func InsertLL(ob interface{}, current *Node) *Node {
	node := &Node{
		OB:    ob,
		Previ: current,
	}

	// current.Next = ll

	return node
}

// PrintLL print linklist
func PrintLL(current Node) {
	for current.Previ != nil {
		fmt.Println(current.OB)
		current = *current.Previ
	}
}
