package Assignment

func SearchElement(nums []int, B int) bool {

	for i := range nums {
		if nums[i] == B {
			return true
		}
	}
	return false

}