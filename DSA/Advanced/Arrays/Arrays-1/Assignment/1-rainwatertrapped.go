package Assignment

func trap(A []int) int {
	ans := 0
	leftMaxArray := getLeftMaxArray(A)
	rightMaxArray := getRightMaxArray(A)
	for i := range A {
		temp := min(leftMaxArray[i], rightMaxArray[i]) - A[i]
		if temp > 0 {
			ans += temp
		}
	}
	return ans
}

func getLeftMaxArray(A []int) []int {

	leftMaxArray := make([]int, len(A))
	leftMaxArray[0] = A[0]
	for i := 1; i < len(A); i++ {
		leftMaxArray[i] = maximum(leftMaxArray[i-1], A[i])
	}

	return leftMaxArray

}

func getRightMaxArray(A []int) []int {

	rightMaxArray := make([]int, len(A))
	rightMaxArray[len(A)-1] = A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		rightMaxArray[i] = maximum(rightMaxArray[i+1], A[i])
	}

	return rightMaxArray

}

func maximum(a, b int) int {
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
