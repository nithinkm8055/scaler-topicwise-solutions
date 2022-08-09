package Assignment

func MaxArea(A []int) int {
	i := 0
	j := len(A) - 1

	ans := 0
	for i < j {
		ans = max(ans, (j-i)*min(A[i], A[j]))

		if A[i] < A[j] {
			i++
		} else {
			j--
		}
	}
	return ans
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
