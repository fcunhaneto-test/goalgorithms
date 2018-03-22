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
type LL = *Node

// head are head of link list
var head = &Node{nil, nil, nil}

// tail are tail of link list
var tail = &Node{nil, nil, nil}

func init() {
	head.N = nil
	head.Next = tail
	head.Previ = nil

	tail.N = nil
	tail.Next = nil
	tail.Previ = head
}

// LlInit init list link
func LlInit(v interface{}) LL {
	var node Node
	node.N = v
	node.Next = tail
	node.Previ = head

	head.Next = &node
	tail.Previ = &node

	return &node
}

// LlEnqueue insert node in tail of link list
func LlEnqueue(v interface{}, current LL) LL {
	var node Node

	if current == nil || LlEmpty() {
		return LlInit(v)
	}

	node.N = v
	node.Next = tail
	node.Previ = current
	current.Next = &node
	tail.Previ = &node

	return &node
}

// LlPush insert node in head of link list
func LlPush(v interface{}, current LL) LL {
	var node Node

	if current == nil || LlEmpty() {
		return LlInit(v)
	}

	node.N = v
	node.Next = current
	node.Previ = head

	head.Next = &node
	current.Previ = &node

	return &node
}

// LlInsertBefore insert node before current node
func LlInsertBefore(v interface{}, current LL) {
	var node Node

	node.N = v
	node.Next = current
	node.Previ = current.Previ

	current.Previ = &node
}

// LlInsertAfter insert node after current node
func LlInsertAfter(v interface{}, current LL) {
	var node Node

	node.N = v
	node.Next = current.Next
	node.Previ = current

	current.Next = &node
}

// GetHead return the head of link lists
func GetHead() LL {
	return head.Next
}

// GetTail return the tail of link lists
func GetTail() LL {
	return tail.Previ
}

// LlDeleteNode delete node in link list
func LlDeleteNode(n LL) {
	a := n.Next
	b := n.Previ
	a.Previ = b
	b.Next = a

	n.N = nil
	n.Next = nil
	n.Previ = nil
}

// LlFindByS find node in link list by string value
func LlFindByS(n string, field int) LL {
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
func LlFindByI(i int, field int) LL {
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

func LlEmpty() bool {
	if head.Next == tail {
		return true
	}

	return false
}

// LlPrint print link list
func LlPrint() {
	var node LL
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
	var node LL
	node = tail.Previ

	fmt.Println(node.N)
	for node.Previ != nil {
		node = node.Previ

		if node.Previ != nil {
			fmt.Println(node.N)
		}
	}
}
