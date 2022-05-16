package Assignment

import (
	"math"
)

func PickfromBothSides(nums []int, B int) int {

	max := math.MinInt

	prefixSumSlice := make([]int, B)
	suffixSumSlice := make([]int, B)

	prefixSumSlice[0] = nums[0]
	suffixSumSlice[0] = nums[len(nums)-1]
	for i := 1; i < B; i++ {
		prefixSumSlice[i] = nums[i] + prefixSumSlice[i-1]
	}

	for i := 1; i < B; i++ {
		suffixSumSlice[i] = suffixSumSlice[i-1] + nums[len(nums)-1-i]
	}

	max = maximum(prefixSumSlice[B-1], suffixSumSlice[B-1])

	for i := 0; i < B-1; i++ {
		max = maximum(prefixSumSlice[i]+suffixSumSlice[B-2-i], max)
	}

	return max

}

func maximum(a, b int) int {
	if a > b {
		return a
	}
	return b
}
