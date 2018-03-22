package llQueue

import (
	"fmt"
	"reflect"
	"strings"
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
var Head = &Node{nil, nil, nil}

// tail are tail of link list
var tail = &Node{nil, nil, nil}

func init() {
	Head.V = nil
	Head.Next = tail
	Head.Previ = nil

	tail.V = nil
	tail.Next = nil
	tail.Previ = Head
}

// LlInit init list link
func llInit(v interface{}) Node {
	var node Node
	node.V = v
	node.Next = tail
	node.Previ = Head

	Head.Next = &node
	tail.Previ = &node

	return node
}

// LlPush insert node in the queue
func LlPush(v interface{}) LL {
	var node Node
	current := Head

	if Head.Next == tail {
		node = llInit(v)
		return &node
	}

	for current.Next != nil {
		current = current.Next
	}

	node.V = v
	node.Next = current
	current.Previ.Next = &node
	node.Previ = current.Previ
	tail.Previ = &node

	return &node
}

// LlPop remove node from queue
func LlPop() interface{} {
	if LlEmpty() {
		return nil
	}
	node := *Head.Next
	llDeleteNode(Head.Next)

	return node.V
}

// LlEmpty check if the list is empty
func LlEmpty() bool {
	if Head.Next == tail {
		return true
	}

	return false
}

// LlDeleteNode delete node in link list
func llDeleteNode(node LL) {
	if node.Previ == Head {
		Head.Next = node.Next
		node.Next.Previ = Head
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
func LlFind(n string) LL {
	var s string
	node := Head.Next

	for node.Next != nil {
		s = string(reflect.ValueOf(node.V).Field(0).String())
		if strings.Compare(s, n) == 0 {
			return node
		}

		node = node.Next
	}
	return nil
}

// LlPrint print link list
func LlPrint() {
	var node LL
	node = Head.Next
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
