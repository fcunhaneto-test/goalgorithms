package main

import (
	"fmt"
	"goalgorithms/linklist"
)

// Person type person
type Person struct {
	ID   int
	Name string
}

func main() {
	var s string
	var p Person
	var head, tail, current linklist.LL

	p.ID = 1
	p.Name = "Francisco"
	head = linklist.LlInit(p)

	p.ID = 2
	p.Name = "Erica"
	tail = linklist.LlPush(p, head)

	p.ID = 3
	p.Name = "João"
	tail = linklist.LlPush(p, tail)

	p.ID = 4
	p.Name = "Paula"
	tail = linklist.LlPush(p, tail)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println("Enter person name for delete:")
	fmt.Scanf("%s", &s)

	current = linklist.LlFindByS(s, 2)
	fmt.Println("Find:", current)
}
