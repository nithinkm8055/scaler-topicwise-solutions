package Homework

func books(A []int, B int) int {

	if B > len(A) {
		return -1
	}

	l := A[0]
	r := 0

	// define the answer space
	// min range is book with max pages, and max is sum of all pages of all books
	for i := range A {
		if A[i] > l {
			l = A[i]
		}
		r += A[i]
	}

	ans := -1

	for l <= r {

		mid := (l + r) / 2

		if check(A, mid, B) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans
}

// target is the number of pages
// check checks if target number of pages can be divided among numStudents
func check(A []int, targetPages, numStudents int) bool {

	numStudents = numStudents - 1
	sum := 0
	for i := range A {

		sum += A[i]
		if sum > targetPages {
			sum = A[i]
			numStudents--
		}
		if numStudents < 0 {
			return false
		}
	}

	return true
}
