package Homework

import "sort"

func RepeatedNumber(A []int) []int {
	sort.Ints(A)
	sum := A[0]
	double := 0
	for i := 1; i < len(A); i++ {
		sum += A[i]
		if A[i]-A[i-1] == 0 {
			double = A[i]
		}
	}

	givenSum := (len(A) * (len(A) + 1)) / 2

	return []int{double, (givenSum - sum) + double}
}
