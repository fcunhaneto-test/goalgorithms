package sort

/*
Quicksort ordering algorithm Quick Sort
*/
func Quicksort(a []int, ini int, end int) {
	var pivo int
	if end > ini {
		pivo = partition(a, ini, end)
		Quicksort(a, ini, pivo-1)
		Quicksort(a, pivo+1, end)
	}
}

func partition(a []int, ini int, end int) int {
	var aux int
	l := ini
	r := end
	pivo := a[ini]
	for l < r {
		for a[l] <= pivo {
			l++
		}
		for a[r] > pivo {
			r--
		}
		if l < r {
			aux = a[l]
			a[l] = a[r]
			a[r] = aux
		}
	}
	a[ini] = a[r]
	a[r] = pivo
	return r
}
