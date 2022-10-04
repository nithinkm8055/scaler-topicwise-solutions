package Homework

func BClosestPointsToOrigin(A [][]int, B int) [][]int {

	for i := len(A)/2 - 1; i >= 0; i-- {
		heapifyPoints(A, i)
	}

	size := len(A)
	result := make([][]int, 0)

	for size > 0 {

		if len(result) == B {
			return result
		}

		result = append(result, A[0])
		A[0], A[size-1] = A[size-1], A[0]
		size--
		heapifyPoints(A[:size], 0)
	}

	return result
}

func heapifyPoints(heap [][]int, i int) {

	for 2*i+1 < len(heap) {
		y := 2 * i
		z := []int{1000000, 1000000}

		if y+2 < len(heap) {
			z = heap[y+2]
		}
		x := minSlice(heap[i], heap[y+1], z)

		if x[0] == heap[i][0] && x[1] == heap[i][1] {
			return
		} else if x[0] == heap[y+1][0] && x[1] == heap[y+1][1] {
			heap[i], heap[y+1] = heap[y+1], heap[i]
			i = y + 1
		} else {
			heap[i], heap[y+2] = heap[y+2], heap[i]
			i = y + 2
		}
	}
}

func minSlice(a []int, b []int, c []int) []int {

	asize := pow(a[0], 2) + pow(a[1], 2)
	bsize := pow(b[0], 2) + pow(b[1], 2)
	csize := pow(c[0], 2) + pow(c[1], 2)

	if asize <= bsize && asize <= csize {
		return a
	} else if bsize <= csize && bsize <= asize {
		return b
	}
	return c
}

func pow(a, b int) int {
	if b == 0 {
		return 1
	}
	x := pow(a, b/2)
	if b%2 == 0 {
		return x * x
	}
	return a * x * x
}
