package Homework

func SearchInBitonicArray(A []int, B int) int {

	// check Pivot point, could also be not the pivot point, if B was already found
	pivotIdx, found := checkPivot(A, B)
	if found {
		return pivotIdx
	}

	idx := checkLeft(A[0:pivotIdx], B)
	if idx != -1 {
		return idx
	}

	idx1 := checkRight(A[pivotIdx:], B)

	if idx1 != -1 {
		return idx1 + pivotIdx
	}
	return -1
}

func checkLeft(A []int, B int) int {
	l := 0
	r := len(A) - 1

	for l <= r {
		mid := (l + r) / 2
		if A[mid] == B {
			return mid
		}

		if A[mid] > B {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

func checkRight(A []int, B int) int {
	l := 0
	r := len(A) - 1

	for l <= r {
		mid := (l + r) / 2
		if A[mid] == B {
			return mid
		}

		if A[mid] < B {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return -1
}

// return pivot index
func checkPivot(A []int, B int) (int, bool) {

	l := 0
	r := len(A) - 1

	for l <= r {

		mid := (l + r) / 2

		if A[mid] == B {
			return mid, true
		}

		if A[mid] > A[mid-1] && A[mid] > A[mid+1] {
			return mid, false
		} else if A[mid] > A[mid-1] && A[mid] < A[mid+1] {
			l = mid + 1
		} else if A[mid] < A[mid-1] && A[mid] > A[mid+1] {
			r = mid - 1
		}
	}

	return 0, false
}
