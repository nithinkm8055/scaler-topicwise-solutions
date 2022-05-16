package Homework

func TimetoEquality(A []int) int {

	max := 0
	sum := 0
	for i := range A {
		sum += A[i]
		if A[i] > max {
			max = A[i]
		}
	}

	totalSumAfterChange := max * len(A)
	return totalSumAfterChange - sum
}
