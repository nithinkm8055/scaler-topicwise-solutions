package Assignment

func Reverse(A []int) []int {

	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-1-i] = A[len(A)-1-i], A[i]
	}

	return A
}
