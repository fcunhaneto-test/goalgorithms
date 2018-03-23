package llqueue

import "fmt"

// Node struct for link list
type Node struct {
	N     interface{}
	Next  *Node
	Previ *Node
}

// LL alias for pointer to node
type LL = *Node

// I interface for link list
type I interface {
	LlPush(v interface{})
	LlPop()
}

// head are head of link list
var head = &Node{nil, nil, nil}

// tail are tail of link list
var tail = &Node{nil, nil, nil}

// LlStart start link list
func LlStart() LL {
	head.N = nil
	head.Next = tail
	head.Previ = nil

	tail.N = nil
	tail.Next = nil
	tail.Previ = head

	return head
}

// LlInit input first node in list link
func llInit(v interface{}) LL {
	var node Node

	node.N = v
	node.Next = tail
	node.Previ = head

	head.Next = &node
	tail.Previ = &node

	return &node
}

// LlPush insert node in head of link list
func (current LL) LlPush(v interface{}) LL {
	var node Node

	if current == nil || LlEmpty() {
		return llInit(v)
	}

	node.N = v
	node.Next = current
	node.Previ = head

	head.Next = &node
	current.Previ = &node

	return &node
}

// LlPop remove node from queue
func (current LL) LlPop() (interface{}, LL) {
	if LlEmpty() {
		return nil, nil
	}
	node := *head.Next
	llDeleteNode(head.Next)

	return node.N, head.Next
}

// LlDeleteNode delete node in link list
func llDeleteNode(n *Node) {
	a := n.Next
	b := n.Previ
	a.Previ = b
	b.Next = a

	n.N = nil
	n.Next = nil
	n.Previ = nil
}

// LlEmpty verify if link list is empty
func LlEmpty() bool {
	if head.Next == tail {
		return true
	}

	return false
}

// LlPrint print link list
func LlPrint() {
	var node *Node
	node = head.Next
	fmt.Println(node.N)
	for node.Next != nil {
		node = node.Next
		if node.Next != nil {
			fmt.Println(node.N)
		}
	}
}
