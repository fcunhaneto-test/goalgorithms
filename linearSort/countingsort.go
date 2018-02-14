package linearSort

// Countingsort sort array by counting ordering
func Countingsort(a []int, k int) []int {
	n1 := len(a)
	n2 := k + 1
	b := make([]int, n1)
	c := make([]int, n2)

	for i := 0; i < n1; i++ {
		c[a[i]] = c[a[i]] + 1
	}

	for i := 1; i < n2; i++ {
		c[i] += c[i-1]
	}

	for i := (n1 - 1); i >= 0; i-- {
		b[c[a[i]]-1] = a[i]
		c[a[i]]--
	}

	return b
}
