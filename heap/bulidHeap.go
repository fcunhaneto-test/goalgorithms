package heap

// BuildHeap construt and maxify a binary tree
func BuildHeap(a []int, n int) {
	for i := n / 2; i > -1; i-- {
		MaintainHeapfy(a, i, n)
	}
}
