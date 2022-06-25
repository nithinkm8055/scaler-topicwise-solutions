package Assignment

func Maxmin(nums []int) (int, int) {

	max := nums[0]
	min := nums[0]

	for i := range nums {
		if nums[i] > max {
			max = nums[i]
		} else if nums[i] < min {
			min = nums[i]
		}
	}

	return max, min

}
