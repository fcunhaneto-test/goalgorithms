package arrayMax

import (
	"math"
)

// SubArrayMax use Kadane algorithm to find the indexes and the value of the largest subarray.
func SubArrayMax(a []int) (int, int, int) {
	x, y := 0, 0
	maxNow := 0
	maxTotal := math.MinInt64
	n := len(a)
	for i := 0; i < n; i++ {
		maxNow = maxNow + a[i]
		if maxNow < 0 {
			maxNow = 0
			x = i + 1
		}
		if maxNow > maxTotal {
			maxTotal = maxNow
			y = i
		}
	}
	return x, y, maxTotal
}
