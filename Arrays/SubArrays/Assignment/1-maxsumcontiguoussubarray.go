package Assignment

func MaxSubArray(nums []int) int {

	currSum := 0
	maxSum := nums[0]
	for i := 1; i < len(nums); i++ {
		currSum += nums[i]
		if currSum > maxSum {
			maxSum = currSum
		}
		if currSum < 0 { // if all negatives
			currSum = 0
		}
	}
	return maxSum
}
