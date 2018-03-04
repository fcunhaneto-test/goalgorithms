package linklist

import "fmt"

// Node struct for link list
type Node struct {
	V     interface{}
	Next  *Node
	Previ *Node
}

// LL link list
type LL = *Node

// InitLL init list link
func InitLL() LL {
	var head Node
	head.V = nil
	head.Next = nil
	head.Previ = nil

	return &head
}

// InsertLL insert struct in link list
func InsertLL(v interface{}, current LL) LL {
	var node Node
	node.V = v
	node.Next = nil
	node.Previ = current

	current.Next = &node
	return &node
}

// GetNext get next node
func GetNext(current LL) LL {
	return current.Next
}

// PrintLL print link list
func PrintLL(head LL) {
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.V)
	}
}

// PrintLLReverse print link list in reverse order
func PrintLLReverse(current LL) {
	fmt.Println(current.V)
	for current.Previ != nil {
		current = current.Previ
		if current.V != nil {
			fmt.Println(current.V)
		}
	}
}
