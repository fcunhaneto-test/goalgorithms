package sort

/*
Insertsort ordering algorithm Insertion Sort
*/
func Insertsort(a []int, n int) {
	var i, j, key int
	for j = 1; j < n; j++ {
		key = a[j]
		i = j - 1
		for (i >= 0) && (a[i] > key) {
			a[i+1] = a[i]
			i = i - 1
		}
		a[i+1] = key
	}

	Printmat(a, n)
}
