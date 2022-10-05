package Assignment

func AthLargestElement(A int, B []int) []int {

	// build a min heap of size A

	for i := A/2 - 1; i >= 0; i-- {
		heapify(B[:A], i)
	}

	result := make([]int, 0)

	for i := 0; i < A-1; i++ {
		result = append(result, -1)
	}

	for i := A; i < len(B); i++ {
		result = append(result, B[0])

		if B[i] >= B[0] {
			B[0], B[i] = B[i], B[0]
			heapify(B[:A], 0)
		}

	}
	result = append(result, B[0])

	return result
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
