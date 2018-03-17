package main

import (
	"fmt"
	"goalgorithms/llQueue"
)

// Person type person
type Person struct {
	ID   int
	Name string
}

func main() {
	var s string
	var p Person
	var node interface{}
	var current llQueue.LL

	p.ID = 1
	p.Name = "Francisco"
	llQueue.LlPush(p)

	p.ID = 2
	p.Name = "Erica"
	llQueue.LlPush(p)

	p.ID = 3
	p.Name = "João"
	llQueue.LlPush(p)

	p.ID = 4
	p.Name = "Paula"
	llQueue.LlPush(p)

	fmt.Println()
	llQueue.LlPrint()
	fmt.Println()
	llQueue.LlPrintReverse()
	fmt.Println("Type enter to continue:")
	fmt.Scanf("%s", &s)

	current = llQueue.LlFind(3)
	fmt.Println("Find 3:", current)

	node = llQueue.LlPop()
	fmt.Println(node.(Person))

	p.ID = 5
	p.Name = "Maria"
	llQueue.LlPush(p)

	for true {
		node = llQueue.LlPop()
		if llQueue.LlEmpty() {
			break
		}
		fmt.Println(node.(Person))
	}

}
