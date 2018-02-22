package ex_linklist

import "fmt"

type Person struct {
	ID    int
	Next  *Person
	Previ *Person
}

func InsertP(id int, current Person) Person {
	var n Person
	n.ID = id
	current.Next = &n
	n.Previ = &current

	return n
}

func PrintP(current Person) {

	for current.Previ != nil {
		fmt.Println(current.ID)
		current = *current.Previ
	}
}
