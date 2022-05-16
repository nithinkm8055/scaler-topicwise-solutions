package Homework

func LeftRotate(A []int, B []int) [][]int {

	result := make([][]int, 0)

	for i := 0; i < len(B); i++ {
		var temp []int
		temp = append(temp, A...)
		if B[i] > len(A) {
			B[i] = B[i] % len(A)
		}

		reverse(temp)
		reverse(temp[:len(A)-B[i]])
		reverse(temp[len(A)-B[i]:])

		result = append(result, temp)
	}

	return result

}

func reverse(A []int) {
	for i := 0; i < len(A)/2; i++ {
		A[i], A[len(A)-i-1] = A[len(A)-i-1], A[i]
	}
}
