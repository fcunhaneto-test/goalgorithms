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
	var current *linklist.Node

	current = linklist.LlStart()
	fmt.Println("current:", current)

	p.ID = 1
	p.Name = "Francisco"
	current = current.LlPush(p)
	fmt.Println("first:", current)

	p.ID = 2
	p.Name = "Erica"
	current = current.LlPush(p)

	p.ID = 3
	p.Name = "João"
	current = current.LlPush(p)

	p.ID = 4
	p.Name = "Paula"
	current = current.LlPush(p)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println("Enter person name for delete:")
	fmt.Scanf("%s", &s)

	// current.LL = linklist.LlFindByS(s, 2)
	// fmt.Println("Find:", current)
}
