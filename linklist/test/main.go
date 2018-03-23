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
	var head, tail, current linklist.LList

	p.ID = 1
	p.Name = "Francisco"
	head.LlInit(p)

	p.ID = 2
	p.Name = "Erica"
	tail.LlPush(p)

	p.ID = 3
	p.Name = "João"
	tail.LlPush(p)

	p.ID = 4
	p.Name = "Paula"
	tail.LlPush(p)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println("Enter person name for delete:")
	fmt.Scanf("%s", &s)

	current.LL = linklist.LlFindByS(s, 2)
	fmt.Println("Find:", current)
}
