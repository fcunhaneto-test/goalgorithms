package mdc

// Calculate gcd with Euclid's algorithm
func MdcEuclides(a int, b int) int {

	r1 := b % a

	if r1 == 0 {
		return a
	}

	return MdcEuclides(r1, a)

}
