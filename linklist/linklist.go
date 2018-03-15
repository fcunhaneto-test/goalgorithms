package linklist

import (
	"fmt"
	"reflect"
)

// Node struct for link list
type Node struct {
	V     interface{}
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

// LlInsertTail insert node in tail of link list
func LlInsertTail(v interface{}, current LL) LL {
	var node Node

	node.V = v
	node.Next = tail
	node.Previ = current
	current.Next = &node
	tail.Previ = &node

	return &node
}

// GetHead return the head of link lists
func GetHead() LL {
	return head.Next
}

// GetTail return the tail of link lists
func GetTail() LL {
	return tail.Previ
}

// LlInsertBefore insert node before current node
func LlInsertBefore(v interface{}, current LL) LL {
	var node Node

	node.V = v
	if current.Previ == head {
		node.Next = current
		node.Previ = head

		head.Next = &node
		current.Previ = &node
	} else {
		node.Next = current
		node.Previ = current.Previ

		current.Previ.Next = &node
		current.Previ = &node
	}

	return &node
}

// LlDeleteNode delete node in link list
func LlDeleteNode(node LL) {
	if node.Previ == head {
		head.Next = node.Next
		node.Next.Previ = head
	} else if node.Next == tail {
		tail.Previ = node.Previ
		node.Previ.Next = tail
	} else {
		// previ := node.Previ
		// next := node.Next

		node.Previ.Next = node.Next
		node.Next.Previ = node.Previ
	}

	node.V = nil
	node.Next = nil
	node.Previ = nil
}

// LlFind find node in link list
func LlFind(n int) LL {
	var num int
	node := head.Next

	for node.Next != nil {
		num = int(reflect.ValueOf(node.V).Field(0).Int())
		if num == n {
			return node
		}

		node = node.Next
	}
	return nil
}

// LlPrint print link list
func LlPrint() {
	var node LL
	node = head.Next
	fmt.Println(node.V)
	for node.Next != nil {
		node = node.Next
		if node.Next != nil {
			fmt.Println(node.V)
		}
	}
}

// LlPrintReverse print link list in reverse order
func LlPrintReverse() {
	var node LL
	node = tail.Previ

	fmt.Println(node.V)
	for node.Previ != nil {
		node = node.Previ

		if node.Previ != nil {
			fmt.Println(node.V)
		}
	}
}
