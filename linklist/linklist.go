/*
Package linklist - link list can use as stack or queue
*/
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

// LL alias for pointer to node
type LL = *Node

// I interface for link list
type I interface {
	LlEnqueue(v interface{})
	LlPush(v interface{})
	LlInsertBefore(v interface{})
	LlInsertAfter(v interface{})
	GetHead()
	GetTail()
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

// LlEnqueue insert node in tail of link list
func (current LL) LlEnqueue(v interface{}) LL {
	var node Node

	if current == nil || LlEmpty() {
		return llInit(v)
	}

	node.N = v
	node.Next = tail
	node.Previ = current
	current.Next = &node
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

// LlInsertBefore insert node before current node
func (current LL) LlInsertBefore(v interface{}) {
	var node Node

	node.N = v
	node.Next = current
	node.Previ = current.Previ
	current.Previ = &node
}

// LlInsertAfter insert node after current node
func (current LL) LlInsertAfter(v interface{}) {
	var node Node

	node.N = v
	node.Next = current.Next
	node.Previ = current

	current.Next = &node
}

// GetHead return the head of link lists
func (current LL) GetHead() LL {
	return head.Next
}

// GetTail return the tail of link lists
func (current LL) GetTail() LL {
	return tail.Previ
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
			fmt.Println(node.N, *node)
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
