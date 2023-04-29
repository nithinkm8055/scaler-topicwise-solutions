package Assignment

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
func sortquick(A []int) []int {
	return quickSort(A, 0, len(A)-1)
}

func quickSort(A []int, l, r int) []int {
	if l >= r {
		return A
	}

	index := rearrange(A, l, r)
	quickSort(A, l, index-1)
	quickSort(A, index+1, r)

	return A
}

func rearrange(A []int, l, r int) int {
	p1 := l + 1
	p2 := r

	for p1 <= p2 {

		if A[p1] <= A[l] {
			p1++
		} else if A[p2] > A[l] {
			p2--
		} else {
			A[p1], A[p2] = A[p2], A[p1]
			p1++
			p2--
		}
	}

	A[l], A[p2] = A[p2], A[l]
	return p2
}
