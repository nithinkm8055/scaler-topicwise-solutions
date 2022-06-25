package Sort

// parse repeatedly along the slice and select the minimum element
// from unsorted part and putting it at the beginning

// TC O(n^2)
// TC O(1)
func selectionSort(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		min := 10000000
		index := -1
		for j := i; j < len(nums); j++ {
			if nums[j] < min {
				min = nums[j]
				index = j
			}
		}
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}

func minimum(a, b int) int {
	if a < b {
		return a
	}
	return b
}
