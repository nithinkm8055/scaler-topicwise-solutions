package Assignment

import (
	"math"
)

func Submatrixsumqueries(A [][]int, B []int, C []int, D []int, E []int) []int {
	prefixSum := make([][]int, len(A))
	mod := (int((math.Pow10(9))) + 7)
	// row wise prefixSum
	for i := range A {
		prefixSum[i] = make([]int, len(A[i]))
		for j := 0; j < len(A[i]); j++ {
			if j == 0 {
				prefixSum[i][j] = A[i][j]
			} else {
				prefixSum[i][j] = prefixSum[i][j-1] + A[i][j]
			}

		}
	}

	// colum wise prefixSum
	for j := 0; j < len(A[0]); j++ {
		for i := 0; i < len(A); i++ {
			if i != 0 {
				prefixSum[i][j] = prefixSum[i][j] + prefixSum[i-1][j]
			}
		}
	}
	result := make([]int, 0)

	for i := range B {

		sum := 0

		sum += prefixSum[D[i]-1][E[i]-1]

		if B[i]-1 > 0 {
			sum -= prefixSum[B[i]-2][E[i]-1]
		}
		if C[i]-1 > 0 {
			sum -= prefixSum[D[i]-1][C[i]-2]
		}
		if B[i]-1 > 0 && C[i]-1 > 0 {
			sum = sum + prefixSum[B[i]-2][C[i]-2]
		}
		// prefixSum[D[i]][E[i]] - prefixSum[B[i]-1][E[i]] - prefixSum[D[i]][C[i]-1] + prefixSum[B[i]-1][C[i]-1]
		sum = sum % mod
		result = append(result, sum)
	}
	return result
}
