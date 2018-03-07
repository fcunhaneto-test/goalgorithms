package linklist

import (
	"fmt"
)

// Person type person
type Person struct {
	ID   int
	Name string
}

// Node struct for link list
type Node struct {
	V     interface{}
	Next  *Node
	Previ *Node
}

// LL link list
type LL = *Node

// LlInit init list link
func LlInit(v interface{}) LL {
	var head Node
	head.V = v
	head.Next = nil
	head.Previ = nil

	return &head
}

// LlInsertTail insert struct in tail of link list
func LlInsertTail(v interface{}, current LL) LL {
	var node Node
	node.V = v
	node.Next = nil
	node.Previ = current

	current.Next = &node
	return &node
}

// LlInsertHead insert struct in head of link list
func LlInsertHead(v interface{}, head LL) LL {
	var node Node
	node.V = v
	node.Next = head
	node.Previ = nil

	head.Previ = &node
	return &node
}

// LlInsertAfter insert node after given node
func LlInsertAfter(v interface{}, after LL) {
	var node Node
	node.V = v
	node.Next = after.Next
	node.Previ = after

	after.Next.Previ = &node
	after.Next = &node
}

// LlInsertLlAfter insert link lists after given node
func LlInsertLlAfter(head LL, tail LL, node LL) {
	tail.Next = node.Next
	node.Next.Previ = tail
	node.Next = head
	head.Previ = node
}

// LlDeleteNode delete node in link list
func LlDeleteNode(node LL) {
	previ := node.Previ
	next := node.Next

	previ.Next = next
	next.Previ = previ

	node.V = nil
	node.Next = nil
	node.Previ = nil
}

// LlDeleteHead delete node in link list
func LlDeleteHead(head LL) LL {
	node := head.Next
	node.Previ = nil

	head.V = nil
	head.Next = nil
	head.Previ = nil

	return node
}

// LlFindFromHead find node from the head to tail
func LlFindFromHead(n int, head LL) LL {
	var node Person

	node = head.V.(Person)
	if n == node.ID {
		return head
	}

	for head.Next != nil {
		head = head.Next
		node = head.V.(Person)

		if n == node.ID {
			return head
		}
	}
	return nil
}

// LlFindFromTail find node from the tail to head
func LlFindFromTail(n int, tail LL) LL {
	var node Person

	node = tail.V.(Person)
	if n == node.ID {
		return tail
	}

	for tail.Previ != nil {
		tail = tail.Previ
		node = tail.V.(Person)
		if n == node.ID {
			return tail
		}
	}
	return nil
}

// LlGetPrevi get previous node
func LlGetPrevi(current LL) LL {
	return current.Previ
}

// LlGetNext get next node
func LlGetNext(current LL) LL {
	return current.Next
}

// LlPrint print link list
func LlPrint(head LL) {
	fmt.Println(head.V)
	for head.Next != nil {
		head = head.Next
		fmt.Println(head.V)
	}
}

// LlPrintReverse print link list in reverse order
func LlPrintReverse(current LL) {
	fmt.Println(current.V)
	for current.Previ != nil {
		current = current.Previ
		if current.V != nil {
			fmt.Println(current.V)
		}
	}
}
