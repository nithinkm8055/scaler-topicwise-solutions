package Assignment

func SpecialInteger(A []int, B int) int {

	l := 0
	r := len(A)

	ans := 0
	for l <= r {

		mid := (l + r) / 2

		if checkWindow(A, mid, B) {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}
	}

	return ans
}

func checkWindow(A []int, windowSize, target int) bool {

	sum := 0

	for i := 0; i < windowSize; i++ {
		sum += A[i]
	}

	if sum > target {
		return false
	} else {

		for i := windowSize; i < len(A); i++ {
			sum -= A[i-windowSize]
			sum += A[i]
			if sum > target {
				return false
			}
		}
	}

	return true
}
