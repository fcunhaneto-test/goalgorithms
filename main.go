package main

import (
	"fmt"
	"goalgorithms/linearSort"
)

func main() {
	// a := []int{5, 2, 8, 7, 4, 6, 1, 3}
	// sort.Mergesort(a, 0, len(a)-1)
	// sort.Printmat(a, len(a))
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

	a := []int{2, 1, 3, 4, 1, 2, 1, 5, 4}
	b := linearSort.Countingsort(a, 5)
	fmt.Println(b)
}
