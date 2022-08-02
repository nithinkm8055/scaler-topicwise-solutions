package Assignment

func search(A []int, B int) int {
	l := 0
	r := len(A) - 1

	for l <= r {

		mid := (l + r) / 2

		if A[mid] == B {
			return mid
		} else if A[mid] >= A[l] {

			if A[mid] > B && B >= A[l] {
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else if A[mid] < A[l] {
			if B > A[mid] && B <= A[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}

	return -1
}
