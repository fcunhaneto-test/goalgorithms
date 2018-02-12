package sort

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
a: 	slice representing heap
i: 	first level of tree, for left use 1 and for right use 2
n: 	array size
*/
func maintainHeapfy(a []int, i int, n int) {
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
		maintainHeapfy(a, max, n)
	}
}

/*
buildHeap create array heap max

param:
a: 	slice representing heap
n: 	array size
*/
func buildHeap(a []int, n int) {
	for i := n / 2; i > -1; i-- {
		maintainHeapfy(a, i, n)
	}
}

/*
HeapSort sort a array heap max

param:
a: 	slice representing heap
n: 	array size
*/
func Heapsort(a []int, n int) {
	buildHeap(a, n)
	for i := n - 1; i > -1; i-- {
		a[0], a[i] = a[i], a[0]
		maintainHeapfy(a, 0, i)
	}
}
