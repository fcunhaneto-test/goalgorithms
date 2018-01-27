package gcd

func GcdEuclids(r1 int, q1 int) int {
	r2 := q1 % r1

	if r2 == 0 {
		return r1
	}

	for r2 != 0 {
		q1 = r1
		r1 = r2
		r2 = q1 % r1
	}

	return r1
}
