package linklist

import (
	"fmt"
	"reflect"
	"strings"
)

// Node struct for link list
type Node struct {
	N     interface{}
	Next  *Node
	Previ *Node
}

// LL link list
type LList struct {
	LL *Node
}

// head are head of link list
var head = &Node{nil, nil, nil}

// tail are tail of link list
var tail = &Node{nil, nil, nil}

// LlInit init list link
func (current LList) LlInit(v interface{}) {
	var node Node

	head.N = nil
	head.Next = tail
	head.Previ = nil

	tail.N = nil
	tail.Next = nil
	tail.Previ = head

	node.N = v
	node.Next = tail
	node.Previ = head

	head.Next = &node
	tail.Previ = &node

	current.LL = &node
}

// LlEnqueue insert node in tail of link list
func (current LList) LlEnqueue(v interface{}) {
	var node Node

	node.N = v
	node.Next = tail
	node.Previ = current.LL
	current.LL.Next = &node
	tail.Previ = &node

	current.LL = &node
}

// LlPush insert node in head of link list
func (current LList) LlPush(v interface{}) {
	var node Node

	node.N = v
	node.Next = current.LL
	node.Previ = head

	head.Next = &node
	current.LL.Previ = &node

	current.LL = &node
}

// LlInsertBefore insert node before current node
func (current LList) LlInsertBefore(v interface{}) {
	var node Node

	node.N = v
	node.Next = current.LL
	node.Previ = current.LL.Previ

	current.LL.Previ = &node
}

// LlInsertAfter insert node after current node
func (current LList) LlInsertAfter(v interface{}) {
	var node Node

	node.N = v
	node.Next = current.LL.Next
	node.Previ = current.LL

	current.LL.Next = &node
}

// GetHead return the head of link lists
func GetHead() *Node {
	return head.Next
}

// GetTail return the tail of link lists
func GetTail() *Node {
	return tail.Previ
}

// LlDeleteNode delete node in link list
func LlDeleteNode(n *Node) {
	a := n.Next
	b := n.Previ
	a.Previ = b
	b.Next = a

	n.N = nil
	n.Next = nil
	n.Previ = nil
}

// LlFindByS find node in link list by string value
func LlFindByS(n string, field int) *Node {
	var s string
	node := head.Next
	f := field - 1

	for node.Next != nil {
		s = string(reflect.ValueOf(node.N).Field(f).String())
		if strings.Compare(s, n) == 0 {
			return node
		}

		node = node.Next
	}

	return nil
}

// LlFindByI find node in link list by int value
func LlFindByI(i int, field int) *Node {
	var n int
	node := head.Next
	f := field - 1

	for node.Next != nil {
		n = int(reflect.ValueOf(node.N).Field(f).Int())
		if n == i {
			return node
		}
		fmt.Println("Node:", node)
		node = node.Next
	}

	return nil
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

// LlPrintReverse print link list in reverse order
func LlPrintReverse() {
	var node *Node
	node = tail.Previ

	fmt.Println(node.N)
	for node.Previ != nil {
		node = node.Previ

		if node.Previ != nil {
			fmt.Println(node.N)
		}
	}
}
