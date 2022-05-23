package Assignment

import "sort"

func ArithmeticProgression(A []int) int {
	sort.Ints(A)

	d := A[1] - A[0]

	for i := 2; i < len(A); i++ {
		if A[i]-A[i-1] != d {
			return 0
		}
	}

	return 1

}
