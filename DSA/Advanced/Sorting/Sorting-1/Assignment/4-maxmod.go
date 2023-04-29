package Assignment

func MaxMod(A []int) int {

	max := -1

	for i := 0; i < len(A)-1; i++ {
		for j := 0; j < len(A); j++ {
			if A[j] != 0 && A[i]%A[j] > max {
				max = A[i] % A[j]
			}
		}
	}

	return max
}
