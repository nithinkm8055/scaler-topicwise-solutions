package Homework

// n^3 approach
func MaxSumSubMatrix(A [][]int) int {
	pfColSum := make([][]int, len(A))

	for i := 0; i < len(A); i++ {
		pfColSum[i] = make([]int, len(A[i]))
	}

	for j := 0; j < len(A[0]); j++ {

		for i := 0; i < len(A); i++ {
			if i == 0 {
				pfColSum[i][j] = A[i][j]
			} else {
				pfColSum[i][j] = pfColSum[i-1][j] + A[i][j]
			}
		}
	}

	ans := -10000000000
	for tr := 0; tr < len(A); tr++ {

		for br := tr; br < len(A); br++ {

			sum := 0
			for col := 0; col < len(A[tr]); col++ {
				if tr == 0 {
					sum += pfColSum[br][col]
				} else {
					sum += pfColSum[br][col] - pfColSum[tr-1][col]
				}
				if sum > ans {
					ans = sum
				}
				if sum < 0 {
					sum = 0
				}
			}

		}
	}

	return ans
}

// TLE: TC(O(m^2n^2))
func MaxSumSubMatrix2(A [][]int) int {

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
