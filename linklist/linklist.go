package linklist

import "fmt"

// Node struct for link list
type Node struct {
	OB    interface{}
	Next  *Node
	Previ *Node
}

// LLInit init list link
func LLInit(v interface{}) Node {
	var root Node

	root.OB = v
	root.Next = nil
	root.Previ = nil

	return root
}

// InsertLL insert struct in link list
func InsertLL(ob interface{}, current Node) Node {
	var ll Node
	ll.OB = ob
	current.Next = &ll
	ll.Previ = &current

	return ll
}

// PrintLL print linklist
func PrintLL(current Node) {
	for current.Previ != nil {
		fmt.Println(current.OB)
		current = *current.Previ
	}
}
