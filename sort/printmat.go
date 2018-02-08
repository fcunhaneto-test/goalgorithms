package sort

import "fmt"

// Printmat print array
func Printmat(a []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
}
