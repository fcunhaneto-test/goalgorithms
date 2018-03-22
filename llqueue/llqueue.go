package llqueue

import (
	"fmt"
)

// Node struct for link list
type Node struct {
	V     interface{}
	Next  *Node
	Previ *Node
}

// LL link list
type LL = *Node

// Head are Head of link list
var head = &Node{nil, nil, nil}

// tail are tail of link list
var tail = &Node{nil, nil, nil}

func init() {
	head.V = nil
	head.Next = tail
	head.Previ = nil

	tail.V = nil
	tail.Next = nil
	tail.Previ = head
}

// LlInit init list link
func LlInit(v interface{}) LL {
	var node Node
	node.V = v
	node.Next = tail
	node.Previ = head

	head.Next = &node
	tail.Previ = &node

	return &node
}

// LlPush insert node in the queue
func LlPush(v interface{}, current LL) LL {
	var node Node

	node.V = v
	node.Next = current
	node.Previ = head
	current.Previ = &node
	head.Next = &node

	return &node
}

// LlPrint print link list
func LlPrint() {
	var node LL
	node = head
	for node.Next != nil {
		node = node.Next
		fmt.Println(node)
	}
}
