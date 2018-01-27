package heap

// HeapSort sort a array heap
func HeapSort(a []int, n int) {
	BuildHeap(a, n)
	for i := n - 1; i > -1; i-- {
		a[0], a[i] = a[i], a[0]
		MaintainHeapfy(a, 0, i)
	}
}
