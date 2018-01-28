package main

import (
	"goalgorithms/heap"
	"rwarray"
)

func main() {
	// a := []int{5, 2, 8, 7, 4, 6, 1, 3}
	// sort.Mergesort(a, 0, len(a)-1)
	// sort.Printmat(a, len(a))
	// mdc := gcd.GcdEuclids(234, 357)
	// fmt.Println(mdc)

	a := []int{15, 13, 2, 25, 7, 17, 20, 8, 4}
	rwarray.PrintArray(a, len(a))
	heap.HeapSort(a, len(a))
	rwarray.PrintArray(a, len(a))
}
