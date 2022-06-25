package Assignment

// return index
func SubarrayLeastAverage(nums []int, B int) int {
	prefixSumSlice := make([]int, len(nums))
	prefixSumSlice[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		prefixSumSlice[i] = nums[i] + prefixSumSlice[i-1]
	}

	min := 10000000000
	rIndex := -1
	for i := 0; i < len(nums)-B+1; i++ {
		sum := 0
		if i-1 < 0 {
			sum = prefixSumSlice[B+i-1]
		} else {
			sum = prefixSumSlice[B+i-1] - prefixSumSlice[i-1]
		}

		if sum < min {
			rIndex = i
			min = sum
		}
	}
	return rIndex
}
