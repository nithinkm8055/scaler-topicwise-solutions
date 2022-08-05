package BinarySearch_3

func PaintersPartition(A int, B int, C []int) int {

	// define answer space
	// the minimum time to paint the board is the max time in array
	// the max time to paint is the sum of all time elements in array

	l := 0
	r := 0

	for i := range C {
		C[i] = C[i] * B
		r += C[i]
		if C[i] > l {
			l = C[i]
		}
	}

	ans := 0
	for l <= r {

		mid := (l + r) / 2

		if isWorkDone(C, mid, A) {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans % 10000003
}

// isWorkDone checks if work can be done in target amount of time by given numPainters
func isWorkDone(C []int, target, numPainters int) bool {
	sum := 0
	numPainters--
	for i := range C {
		sum += C[i]
		if sum > target {
			sum = C[i]
			numPainters--
		}
	}

	if numPainters < 0 {
		return false
	}

	return true
}
