package main

import (
	"fmt"
	"goalgorithms/sort"
)

func main() {
	// a := []int{5, 2, 8, 7, 4, 6, 1, 3}
	// sort.Mergesort(a, 0, len(a)-1)
	// sort.Printmat(a, len(a))
	// mdc := gcd.GcdEuclids(234, 357)
	// fmt.Println(mdc)

	a := []int{15, 13, 2, 25, 7, 17, 20, 8, 4}
	fmt.Println(a)
	sort.Heapsort(a, len(a))
	fmt.Println(a)
}
