package Homework

func CountingSubarrays(nums []int, B int) int {

	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum < B {
				count++
			} else {
				break
			}
		}
	}

	return count
}
