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

	linklist.LlDeleteNode(head)

	p.ID = 2
	p.Name = "Erica"
	tail = linklist.LlInsertTail(p, head)

	p.ID = 3
	p.Name = "João"
	tail = linklist.LlInsertTail(p, tail)

	p.ID = 4
	p.Name = "Paula"
	tail = linklist.LlInsertTail(p, tail)

	fmt.Println()
	linklist.LlPrint()
	fmt.Println()
	linklist.LlPrintReverse()
	fmt.Println("Type enter to continue:")
	fmt.Scanf("%s", &s)

	current = linklist.LlFind(3)
	fmt.Println("Find 3:", current)

	// p.ID = 5
	// p.Name = "Rachel"
	// current = linklist.LlInsertBefore(p, current)
	// // fmt.Printf("%v - %p\n", current.Previ, current.Previ)
	// // fmt.Printf("%v - %p\n", current, current)
	// // fmt.Printf("%v - %p\n", current.Next, current.Next)
	// fmt.Println()
	// linklist.LlPrint()
	// fmt.Println()
	// linklist.LlPrintReverse()
	// fmt.Println("Type enter to continue:")
	// fmt.Scanf("%s", &s)

	// p.ID = 6
	// p.Name = "Tercio"
	// head = linklist.LlInsertBefore(p, head)

	// fmt.Println()
	// linklist.LlPrint()
	// fmt.Println()
	// linklist.LlPrintReverse()
	// fmt.Println("Type enter to continue:")
	// fmt.Scanf("%s", &s)

	// // current = linklist.LlFind(2)
	// // fmt.Println("Find 2:", current)

	// linklist.LlDeleteNode(current)

	// fmt.Println()
	// linklist.LlPrint()
	// fmt.Println()
	// linklist.LlPrintReverse()
	// fmt.Println("Type enter to continue:")
	// fmt.Scanf("%s", &s)

	// linklist.LlDeleteNode(head)
	// head = linklist.GetHead()
	// fmt.Println("Head:", head)

	// fmt.Println()
	// linklist.LlPrint()
	// fmt.Println()
	// linklist.LlPrintReverse()
	// fmt.Println("Type enter to continue:")
	// fmt.Scanf("%s", &s)

	// linklist.LlDeleteNode(tail)

	// fmt.Println()
	// linklist.LlPrint()
	// fmt.Println()
	// linklist.LlPrintReverse()
	// fmt.Println()

	// tail = linklist.GetTail()
	// fmt.Println("Tail:", tail)
}
