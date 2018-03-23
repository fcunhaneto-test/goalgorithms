package main

import (
	"fmt"
	"goalgorithms/llqueue"
)

// Person type person
type Person struct {
	Name string
}

func main() {
	var p Person
	var node llqueue.LL
	var current llqueue.LL
	var i interface{}

	current = llqueue.LlStart()

	p.Name = "Francisco"
	current = current.LlPush(p)

	p.Name = "Erica"
	current = current.LlPush(p)

	p.Name = "João"
	current = current.LlPush(p)

	p.Name = "Paula"
	current = current.LlPush(p)

	fmt.Println()
	fmt.Println("*****************************")

	node = current

	for node.Next != nil {
		for node.Next != nil {
			if node.Next != nil {
				fmt.Println(node.N)
			}
			node = node.Next
		}
	}

	fmt.Println("*****************************")
	i, current = current.LlPop()
	fmt.Println(i)

	fmt.Println()
	fmt.Println("*****************************")
	for current.Next != nil {
		for current.Next != nil {
			if current.Next != nil {
				fmt.Println(current.N)
			}
			current = current.Next
		}
	}

	fmt.Println("*****************************")
	i, current = current.LlPop()
	fmt.Println(i)

	fmt.Println()
	fmt.Println("*****************************")
	for current.Next != nil {
		for current.Next != nil {
			if current.Next != nil {
				fmt.Println(current.N)
			}
			current = current.Next
		}
	}

}
