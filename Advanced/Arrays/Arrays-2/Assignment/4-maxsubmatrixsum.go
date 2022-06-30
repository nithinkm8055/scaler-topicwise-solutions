package Assignment

func MaxSubmatrixSum(A [][]int) int64 {

	pfSum := make([][]int, len(A))

	// bottom row wise : right to left
	for i := len(A) - 1; i >= 0; i-- {
		pfSum[i] = make([]int, len(A[i]))
		for j := len(A[i]) - 1; j >= 0; j-- {

			if j < len(A[i])-1 {
				pfSum[i][j] = A[i][j] + pfSum[i][j+1]
			} else {
				pfSum[i][j] = A[i][j]
			}
		}
	}

	// col wise : bottom to top
	for j := len(A[0]) - 1; j >= 0; j-- {
		for i := len(A) - 1; i >= 0; i-- {

			if i != len(A)-1 {
				pfSum[i][j] = pfSum[i][j] + pfSum[i+1][j]
			}
		}
	}

	// find max pfSum cell
	max := -1000000000000
	for i := range pfSum {
		for j := range pfSum[i] {
			if pfSum[i][j] > max {
				max = pfSum[i][j]
			}
		}
	}

	return int64(max)

}
