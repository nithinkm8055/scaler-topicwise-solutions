package Homework

func MaxSumSubMatrix(A [][]int) int {

	pfSum := make([][]int, len(A))

	// row wise pfSum
	for i := 0; i < len(A); i++ {
		pfSum[i] = make([]int, len(A[i]))
		for j := 0; j < len(A[i]); j++ {
			if j == 0 {
				pfSum[i][j] = A[i][j]
			} else {
				pfSum[i][j] = A[i][j] + pfSum[i][j-1]
			}
		}
	}

	// col wise pfSum
	for j := 0; j < len(A[0]); j++ {
		for i := 0; i < len(A); i++ {
			if i != 0 {
				pfSum[i][j] = pfSum[i][j] + pfSum[i-1][j]
			}
		}
	}

	max := -10000000

	for a1 := 0; a1 < len(A); a1++ {
		for b1 := 0; b1 < len(A[0]); b1++ {

			for a2 := a1; a2 < len(A); a2++ {
				for b2 := b1; b2 < len(A[0]); b2++ {
					sum := pfSum[a2][b2]

					if a1 > 0 {
						sum -= pfSum[a1-1][b2]
					}
					if b1 > 0 {
						sum -= pfSum[a2][b1-1]
					}
					if a1 > 0 && b1 > 0 {
						sum += pfSum[a1-1][b1-1]
					}

					if sum > max {
						max = sum
					}

				}
			}

		}
	}

	return max

}
