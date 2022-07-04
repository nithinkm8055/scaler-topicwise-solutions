package Homework

func maximumGap(A []int) int {

	prefixMin := make([]int, len(A))
	maxDistance := -1000000

	prefixMin[0] = A[0]
	for i := 1; i < len(A); i++ {
		prefixMin[i] = min(prefixMin[i-1], A[i])
	}

	i := len(A) - 1
	j := len(A) - 1
	for i >= 0 && j >= 0 {
		if A[j] >= prefixMin[i] {
			maxDistance = max(maxDistance, j-i)
			i--
		} else {
			j--
		}
	}
	return maxDistance
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
