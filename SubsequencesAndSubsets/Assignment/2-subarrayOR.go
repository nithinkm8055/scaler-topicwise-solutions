package Assignment

import "math"

func solve(A []int) int {

	rSum := 0

	for i := range A {
		sum := 0
		for j := i; j < len(A); j++ {
			for k := i; k <= j; k++ {
				sum |= A[j]
			}
			rSum += sum

		}

	}

	return rSum % (int((math.Pow10(9))) + 7)

}
