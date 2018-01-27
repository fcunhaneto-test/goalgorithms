package sort

/*
Mergesort ordering algorithm Merge Sort
*/
func Mergesort(a []int, ini int, end int) {
	var midi int
	if ini >= end {
		return
	}

	midi = (ini + end) / 2
	Mergesort(a, ini, midi)
	Mergesort(a, midi+1, end)
	merge(a, ini, midi, end)
}

func merge(a []int, ini int, midi int, end int) {
	p1, p2 := ini, midi+1
	end1, end2 := true, true
	n := end - ini + 1
	temp := make([]int, n)
	if !(n == 0) {
		for i := 0; i < n; i++ {
			if end1 && end2 {
				if a[p1] < a[p2] {
					temp[i] = a[p1]
					p1++
				} else {
					temp[i] = a[p2]
					p2++
				}
				if p1 > midi {
					end1 = false
				}
				if p2 > end {
					end2 = false
				}
			} else {
				if end1 {
					temp[i] = a[p1]
					p1++
				} else {
					temp[i] = a[p2]
					p2++
				}
			}
		}
		k := ini
		for j := 0; j < n; j++ {
			a[k] = temp[j]
			k++
		}
	}
}
