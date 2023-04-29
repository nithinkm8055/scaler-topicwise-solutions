package Assignment

func solve(A []int) int {

	rightGcd := make([]int, len(A))
	leftGcd := make([]int, len(A))

	rightGcd[len(A)-1] = A[len(A)-1]
	leftGcd[0] = A[0]

	maxGcd := 0

	for i := len(A) - 2; i >= 0; i-- {
		rightGcd[i] = gcd(rightGcd[i+1], A[i])
	}

	for i := 0; i < len(A); i++ {
		if i == 0 {
			maxGcd = max(maxGcd, rightGcd[i+1])

		} else if i == len(A)-1 {
			leftGcd[i] = gcd(A[i], leftGcd[i-1])
			if leftGcd[i-1] > maxGcd {
				maxGcd = leftGcd[i-1]
			}
		} else {
			leftGcd[i] = gcd(A[i], leftGcd[i-1])
			maxGcd = max(maxGcd, gcd(leftGcd[i-1], rightGcd[i+1]))
		}
	}

	return maxGcd

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
