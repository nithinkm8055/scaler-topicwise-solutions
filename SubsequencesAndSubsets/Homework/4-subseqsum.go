package Homework

func subseqSum(A []int, B int) int {

	for i := 0; i < (1 << len(A)); i++ {
		sum := 0

		for j := 0; j < len(A); j++ {
			if i&(1<<j) != 0 {
				sum += A[j]
			}
		}
		if sum == B {
			return 1
		}
	}
	return 0
}
