package Assignment

/**
 * @input A : Integer array
 *
 * @Output Integer
 */
func solve(A []int) int {

	if len(A) == 1 {
		return A[0]
	}

	if A[0] > A[1] {
		return A[0]
	}

	if A[len(A)-1] > A[len(A)-2] {
		return A[len(A)-1]
	}

	l := 1
	r := len(A) - 2

	for l <= r {

		mid := (l + r) / 2
		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return A[mid]
		} else if A[mid] < A[mid+1] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}

	return -1
}
