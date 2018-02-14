package sort

import (
	"math/rand"
	"time"
)

/*
InitQuicksort initializes the Quick Sort algorithm

param:
a[]: slice of array
end: int array end len(a)
*/
func InitQuicksort(a []int, n int) {
	end := n - 1
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(end)
	a[0], a[i] = a[i], a[0]
	quicksort(a, 0, end)
}

func quicksort(a []int, ini, end int) {
	if end > ini {
		pivo := partition(a, ini, end)
		quicksort(a, ini, pivo-1)
		quicksort(a, pivo+1, end)
	}
}

func partition(a []int, ini int, end int) int {
	pivo := a[end]
	l := ini - 1

	for j := ini; j <= (end - 1); j++ {
		if a[j] <= pivo {
			l++
			a[l], a[j] = a[j], a[l]
		}
	}

	a[l+1], a[end] = a[end], a[l+1]

	return l + 1
}
