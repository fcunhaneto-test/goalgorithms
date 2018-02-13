package arrayMaxs

// ArrayMax return the index and largest value of the array.
func ArrayMax(a []int) (int, int) {
	x := 0
	max := a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
			x = i
		}
	}

	return x, max
}
