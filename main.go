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

	a := []int{1, 7, 19, 14, 7, 16, 15, 14, 4, 11, 10, 13, 20, 4, 8, 17, 7, 17, 5, 2}
	rwarray.PrintArray(a, len(a))
	heap.HeapSort(a, len(a))
	rwarray.PrintArray(a, len(a))
}
