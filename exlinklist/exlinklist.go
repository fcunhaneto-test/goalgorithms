package exlinklist

import "fmt"

// Person struct of person
type Person struct {
	ID    int
	Next  *Person
	Previ *Person
}

// InsertP insert person
func InsertP(id int, current Person) Person {
	var n Person
	n.ID = id
	current.Next = &n
	n.Previ = &current

	return n
}

// PrintP print person
func PrintP(current Person) {

	for current.Previ != nil {
		fmt.Println(current.ID)
		current = *current.Previ
	}
}
