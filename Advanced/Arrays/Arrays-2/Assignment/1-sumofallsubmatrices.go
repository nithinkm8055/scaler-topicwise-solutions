package Assignment

func Sumofallsubmatrices(A [][]int) int {

	sum := 0
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {

			count := (i + 1) * (len(A) - i) * (j + 1) * (len(A[i]) - j)
			sum += (count * A[i][j])

		}
	}

	return sum
}
