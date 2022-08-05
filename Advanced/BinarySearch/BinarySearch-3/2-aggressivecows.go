package BinarySearch_3

/**
 * @input A : Integer array
 * @input B : Integer
 *
 * @Output Integer
 */
import "sort"

func AggressiveCows(A []int, B int) int {
	sort.Ints(A)
	l := 1
	r := A[len(A)-1] - A[0]
	ans := 0

	for l <= r {

		mid := (l + r) / 2

		if check(A, mid, B) {
			l = mid + 1
			ans = mid
		} else {
			r = mid - 1
		}

	}
	return ans
}

func check(A []int, mid, B int) bool {
	B = B - 1
	last_placed := A[0]
	for i := 1; i < len(A); i++ {
		if A[i]-last_placed >= mid {
			B--
			last_placed = A[i]
			if B == 0 {
				return true
			}
		}
	}
	return false
}
