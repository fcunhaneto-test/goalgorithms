package lists

import (
	"errors"
)

// Stack a stack struct
type Stack struct {
	a   []int
	top int
}

/*
InitStack starts the stack

param:
list: slice to start stack
length: int the stack lenght
*/
func InitStack(list []int, length int) Stack {
	l := make([]int, length)
	n := len(list) - 1
	if list != nil {
		for i := 0; i <= n; i++ {
			l[i] = list[i]
		}
	}

	s := Stack{
		a:   l,
		top: n,
	}

	return s
}

/*
Push inserting element in the stack
*/
func (s *Stack) Push(num int) (Stack, error) {
	s.top++
	if s.isStackFull() {
		return *s, errors.New("stack is full")
	}

	s.a[s.top] = num

	return *s, nil

}

/*
Pop the last element inserted in the stack
*/
func (s *Stack) Pop() (int, error) {
	s.top--
	if s.isStackEmpty() {
		return 0, errors.New("stack is empty")
	}

	return s.a[s.top], nil
}

func (s *Stack) isStackEmpty() bool {
	if s.top >= 0 {
		return false
	}

	return true
}

func (s *Stack) isStackFull() bool {
	if s.top < len(s.a) {
		return false
	}
	return true
}
