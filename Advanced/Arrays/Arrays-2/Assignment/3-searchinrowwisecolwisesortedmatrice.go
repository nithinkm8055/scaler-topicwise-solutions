package Assignment

func solve(A [][]int, B int) int {

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[i]); j++ {
			if A[i][j] == B {
				return (i+1)*1009 + (j + 1)
			}
		}
	}
	return -1

}
