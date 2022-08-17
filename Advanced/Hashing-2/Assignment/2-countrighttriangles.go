package Assignment

func CountRightTriangles(A []int, B []int) int {

	xMap := make(map[int]int)
	yMap := make(map[int]int)

	ans := 0

	for i := range A {
		xMap[A[i]] = xMap[A[i]] + 1
		yMap[B[i]] = yMap[B[i]] + 1
	}

	for i := 0; i < len(A); i++ {
		ans += (xMap[A[i]] - 1) * (yMap[B[i]] - 1)
	}

	return ans % (1e9 + 7)

}
