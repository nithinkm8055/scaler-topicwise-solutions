package Homework

func solve(A [][]int, B int) int {

	pfSum := make([][]int, len(A))

	// row wise prefix sum
	for i := range A {
		pfSum[i] = make([]int, len(A[i]))
		for j := range A[i] {

			if j == 0 {
				pfSum[i][j] = A[i][j]
			} else {
				pfSum[i][j] = A[i][j] + pfSum[i][j-1]
			}
		}
	}

	// col wise prefix sum
	for j := 0; j < len(A[0]); j++ {
		for i := 0; i < len(A); i++ {
			if i != 0 {
				pfSum[i][j] = pfSum[i][j] + pfSum[i-1][j]
			}

		}
	}

	// 0 1 --> 2 3

	// for every top left i,j : bottom right index will B+i-1,B+j-1
	max := -10000000000000
	for i := range A {
		for j := range A[i] {
			sum := 0
			if B+i-1 < len(A) && B+j-1 < len(A[i]) {

				sum = pfSum[B+i-1][B+j-1]

				if i > 0 {
					sum -= pfSum[i-1][B+j-1]
				}
				if j > 0 {
					sum -= pfSum[B+i-1][j-1]
				}
				if i > 0 && j > 0 {
					sum += pfSum[i-1][j-1]
				}
				if sum > max {
					max = sum
				}
			}
		}
	}

	return max
}
