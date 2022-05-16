package Homework

func GoodSubarrays(nums []int, B int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if (j-i+1)%2 == 0 && sum < B {
				count++
			} else if (j-i+1)%2 != 0 && sum > B {
				count++
			}
		}
	}
	return count
}
