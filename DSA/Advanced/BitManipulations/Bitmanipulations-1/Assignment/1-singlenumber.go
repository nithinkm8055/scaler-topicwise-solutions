package Assignment

func singleNumber(A []int) int {

	ans := 0

	for i := range A {
		ans = ans ^ A[i]
	}

	return ans
}
