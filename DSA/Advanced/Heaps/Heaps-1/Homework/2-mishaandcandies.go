package Homework

func solve(A []int, B int) int {

	// build heap
	for i := len(A)/2 - 1; i >= 0; i-- {
		heapify(A, i)
	}

	sum := 0
	size := len(A)
	for size > 0 {
		if A[0] > B {
			return sum
		}
		r := A[0] % 2
		x := A[0] / 2
		sum = sum + x

		A[0], A[size-1] = A[size-1], A[0]
		size--
		heapify(A[:size], 0)
		A[0] = A[0] + r + x
		heapify(A[:size], 0)
		// fmt.Println(A[:size])
	}

	return sum

}

func heapify(heap []int, i int) {
	for 2*i+1 < len(heap) {
		y := 2 * i
		z := 1000000000
		if y+2 < len(heap) {
			z = heap[y+2]
		}
		x := min(heap[i], heap[y+1], z)
		if x == heap[i] {
			return
		} else if x == heap[y+1] {
			heap[i], heap[y+1] = heap[y+1], heap[i]
			i = y + 1
		} else {
			heap[i], heap[y+2] = heap[y+2], heap[i]
			i = y + 2
		}
	}
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
