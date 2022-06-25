package Assignment

// Kadane's Algorithm
func maxSubArray(A []int) int {
	sum := 0
	ans := -1001

	for i := range A {
		sum += A[i]
		ans = max(ans, sum)

		if sum < 0 {
			sum = 0
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
