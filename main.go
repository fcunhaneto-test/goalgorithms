package main

import (
	"goalgorithms/ex_linklist"
)

func main() {
	// a := []int{8, 2, 5, 7, 4, 6, 3, 1}
	// a := []int{25, 13, 2, 4, 7, 17, 20, 8, 15}
	// a := []int{-8, 4, 67, 90, 13, 54, 21}
	// fmt.Println(a)
	// sort.InitQuicksort(a, len(a))
	// fmt.Println(a)
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(7))

	// mdc := gcd.GcdEuclids(234, 357)
	// fmt.Println(mdc)

	// a := []int{15, 13, 2, 25, 7, 17, 20, 8, 4}
	// fmt.Println(a)
	// sort.Heapsort(a, len(a))
	// fmt.Println(a)

	// a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	// // a := []int{0, 13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// i, j, max := arrayMax.SubArrayMax(a)
	// fmt.Printf("[%d, %d] = %d\n", i, j, max)

	// a := []int{2, 1, 3, 4, 1, 2, 1, 5, 4}
	// b := linearSort.Countingsort(a, 5)
	// fmt.Println(b)

	// s := lists.InitStack(nil, 5)
	// fmt.Println(s)

	// for i := 1; i < 5; i++ {
	// 	s, err := s.Push(i * 10)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(s)
	// 	fmt.Println("Size:", s.Size())
	// }
	// for i := 0; i < 6; i++ {
	// 	n, err := s.Pop()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(n)
	// 	fmt.Println("Size:", s.Size())
	// }

	var start, current ex_linklist.Person
	start.ID = 0
	start.Next = &current
	start.Previ = nil
	current.ID = 0
	current.Next = nil
	current.Previ = &start
	current.Next = nil

	current = ex_linklist.InsertP(1, start)
	current = ex_linklist.InsertP(2, current)
	current = ex_linklist.InsertP(3, current)
	current = ex_linklist.InsertP(4, current)

	ex_linklist.PrintP(current)
}
