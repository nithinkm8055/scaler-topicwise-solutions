package Homework

func reversePairs(A []int) int {
	count := 0
	for i := range A {
		for j := i + 1; j < len(A); j++ {
			if A[i] > 2*A[j] {
				count++
			}
		}
	}
	return count
}
