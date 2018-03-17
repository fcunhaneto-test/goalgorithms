package llQueue

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
func llInit(v interface{}) Node {
	var node Node
	node.V = v
	node.Next = tail
	node.Previ = head

	head.Next = &node
	tail.Previ = &node

	return node
}

// LlPush insert node in the queue
func LlPush(v interface{}) {
	var node Node
	current := head

	if head.Next == tail {
		node = llInit(v)
		return
	}

	for current.Next != nil {
		current = current.Next
	}

	node.V = v
	node.Next = current
	current.Previ.Next = &node
	node.Previ = current.Previ
	tail.Previ = &node
}

// LlPop remove node from queue
func LlPop() interface{} {
	if LlEmpty() {
		return nil
	}
	node := *head.Next
	llDeleteNode(head.Next)

	return node.V
}

// LlEmpty check if the list is empty
func LlEmpty() bool {
	if head.Next == tail {
		return true
	}

	return false
}

// LlDeleteNode delete node in link list
func llDeleteNode(node LL) {
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
