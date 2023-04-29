package Assignment

func searchInsert(A []int, B int) int {
	l := 0
	r := len(A) - 1

	for l <= r {
		mid := (l + r) / 2

		if A[mid] == B {
			return mid
		} else if A[mid] < B {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return l
}
