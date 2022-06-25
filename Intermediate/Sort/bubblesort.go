package Sort

// Bubble sort does in place swapping of 2 elements when they are in wrong order
func bubbleSort(nums []int) []int {

	swapped := true
	for swapped {
		swapped = false

		for i := 1; i < len(nums); i++ {
			if nums[i-1] > nums[i] {
				nums[i], nums[i-1] = nums[i-1], nums[i]
				swapped = true
			}
		}
	}
	return nums
}
