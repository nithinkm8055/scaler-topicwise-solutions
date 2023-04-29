package Homework

import "sort"

func solve(A []int, B int) int {

	sort.Ints(A)
	minSum := 0

	if len(A) > 0 && len(A) > B && B > 0 {
		minSum = A[B-1] - A[0]

		for i := B; i < len(A); i++ {

			if A[i]-A[i-B+1] < minSum {
				minSum = A[i] - A[i-B+1]
			}
		}
	}
	return minSum

}
