package main

import (
	"fmt"
	"goalgorithms/arrayMax"
)

func main() {
	a := []int{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	// a := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	x, y, maxi := arrayMax.SubArrayMax(a)
	fmt.Println(x, y, maxi)
}
