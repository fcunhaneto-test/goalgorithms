package heap

func parent(i int) int {
	return i / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return (2 * i) + 2
}

/*
MaintainHeapfy heap property maintenance

param:
a: 	array representing heap
i: 	first level of tree, for left use 1 and for right use 2
n: 	array size
*/
func MaintainHeapfy(a []int, i int, n int) {
	l := left(i)
	r := right(i)
	max := i

	if l < n && a[l] < a[i] {
		max = l
	}

	if r < n && a[r] < a[max] {
		max = r
	}

	if max != i {
		a[i], a[max] = a[max], a[i]
		MaintainHeapfy(a, max, n)
	}
}
