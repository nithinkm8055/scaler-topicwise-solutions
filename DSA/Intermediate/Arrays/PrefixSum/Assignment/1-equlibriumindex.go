package Assignment

//The equilibrium index of an array is an index such that
//the sum of elements at lower indexes is equal to the sum of elements at higher indexes.

func EquilibriumIndex(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	prefixSumSlice := make([]int, len(nums))

	prefixSumSlice[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		prefixSumSlice[i] = nums[i] + prefixSumSlice[i-1]
	}

	for i := range nums {
		if i > 0 && prefixSumSlice[i-1] == prefixSumSlice[len(nums)-1]-prefixSumSlice[i] {
			return i
		} else if i == 0 && 0 == prefixSumSlice[len(nums)-1]-prefixSumSlice[i] {
			return 0
		}
	}
	return -1
}
