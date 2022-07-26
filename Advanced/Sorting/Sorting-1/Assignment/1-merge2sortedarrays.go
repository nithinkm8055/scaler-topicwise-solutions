package Assignment

func MergeSortedArrays(A []int, B []int) []int {
	i := 0
	j := 0

	C := make([]int, 0)

	for i < len(A) && j < len(B) {

		if A[i] >= B[j] {
			C = append(C, B[j])
			j++
		} else {
			C = append(C, A[i])
			i++
		}

	}
	for i < len(A) {
		C = append(C, A[i])
		i++
	}
	for j < len(B) {
		C = append(C, B[j])
		j++
	}
	return C

}
