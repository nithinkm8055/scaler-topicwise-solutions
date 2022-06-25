package Assignment

// rotate an array B times towards the right
func Rotate(nums []int, rotate int) []int {

	if rotate > len(nums) {
		rotate = rotate % len(nums)
	}

	reverse(nums)
	reverse(nums[:rotate])
	reverse(nums[rotate:])

	return nums
}

func reverse(nums []int) {
	for i := 0; i < len(nums)/2; i++ {
		nums[i], nums[len(nums)-1-i] = nums[len(nums)-1-i], nums[i]
	}
}
