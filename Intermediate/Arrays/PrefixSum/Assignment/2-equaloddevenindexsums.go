package Assignment

//Given an array, arr[] of size N, the task is to find the count of array indices such that
//removing an element from these indices makes the sum of even-indexed and odd-indexed array elements equal.

func EqualOddEvenSum(nums []int) int {

	prefixEvenSlice := make([]int, len(nums))
	prefixOddSlice := make([]int, len(nums))

	prefixEvenSlice[0] = nums[0]
	prefixOddSlice[0] = 0
	for i := 1; i < len(nums); i++ {
		if i%2 == 0 {
			prefixEvenSlice[i] = nums[i] + prefixEvenSlice[i-1]
			prefixOddSlice[i] = prefixOddSlice[i-1]
		} else {
			prefixEvenSlice[i] = prefixEvenSlice[i-1]
			prefixOddSlice[i] = prefixOddSlice[i-1] + nums[i]
		}
	}

	count := 0
	for i := range nums {
		evenSum := 0
		oddSum := 0
		if i == 0 {
			evenSum = prefixOddSlice[len(nums)-1]
			oddSum = prefixEvenSlice[len(nums)-1]
		} else {
			evenSum = prefixEvenSlice[i-1] + (prefixOddSlice[len(nums)-1] - prefixOddSlice[i])
			oddSum = prefixOddSlice[i-1] + (prefixEvenSlice[len(nums)-1] - prefixEvenSlice[i])
		}

		if evenSum == oddSum {
			count++
		}
	}

	return count
}
