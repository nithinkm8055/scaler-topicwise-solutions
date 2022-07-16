package Assignment

func pubg(A []int) int {

	// find the gcd of all numbers
	gcdiv := 0
	for i := range A {
		gcdiv = gcd(gcdiv, A[i])
	}

	return gcdiv
}
