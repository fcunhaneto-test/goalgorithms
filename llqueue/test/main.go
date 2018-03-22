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
	var current llqueue.LL

	p.Name = "Francisco"
	current = llqueue.LlInit(p)

	p.Name = "Erica"
	current = llqueue.LlPush(p, current)

	p.Name = "João"
	current = llqueue.LlPush(p, current)

	p.Name = "Paula"
	current = llqueue.LlPush(p, current)

	fmt.Println()
	llqueue.LlPrint()

}
