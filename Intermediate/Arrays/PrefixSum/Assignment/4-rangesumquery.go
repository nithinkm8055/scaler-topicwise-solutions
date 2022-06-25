package Assignment

func RangeSum(nums []int, B [][]int) []int64 {

	prefixSumSlice := make([]int, len(nums))

	prefixSumSlice[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSumSlice[i] = nums[i] + prefixSumSlice[i-1]
	}

	result := make([]int64, 0)
	for i := range B {
		value := 0
		if B[i][1] == B[i][0] {
			value = nums[B[i][0]-1]
		} else if B[i][0]-2 < 0 {
			value = prefixSumSlice[B[i][1]-1]
		} else {
			value = prefixSumSlice[B[i][1]-1] - prefixSumSlice[B[i][0]-2]
		}
		result = append(result, int64(value))
	}

	return result

}
