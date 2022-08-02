package Assignment

func sqrt(A int) int {

	l := 1
	r := A

	ans := 0

	for l <= r {

		mid := (l + r) / 2

		if mid*mid == A {
			return mid
		} else if mid*mid < A {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return ans

}
