package Sort

// sort all elememts to left
func insertionSort(nums []int) []int {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			nums[i], nums[i+1] = nums[i+1], nums[i]
			for j := i; j >= 1; j-- {
				if nums[j-1] > nums[j] {
					nums[j], nums[j-1] = nums[j-1], nums[j]
				} else {
					break
				}
			}
		}
	}
	return nums
}
