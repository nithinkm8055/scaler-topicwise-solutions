package Assignment

func maxSubarray(A int, B int, C []int) int {
	maxSum := 0
	for i := 0; i < A; i++ {
		sum := 0
		for j := i; j < A; j++ {
			sum += C[j]
			if sum <= B {
				maxSum = maximum(sum, maxSum)
			} else {
				break
			}
		}
	}
	return maxSum
}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
