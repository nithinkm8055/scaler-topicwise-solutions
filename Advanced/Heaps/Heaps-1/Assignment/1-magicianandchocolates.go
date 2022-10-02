package Assignment

func nchoc(A int, B []int) int {

	// build heap
	for i := len(B)/2 - 1; i >= 0; i-- {
		heapify(B, i)
	}

	sum := 0
	for i := 0; i < A; i++ {
		sum += B[0]
		B[0] = B[0] / 2
		heapify(B, 0)
	}

	return sum % (1e9 + 7)

}

func heapify(heap []int, i int) {

	for (2*i)+1 < len(heap) {
		x := max(heap[i], heap[2*i+1], -1)
		if 2*i+2 < len(heap) {
			x = max(heap[i], heap[2*i+1], heap[2*i+2])
		}

		if x == heap[i] {
			return
		} else if x == heap[2*i+1] {
			heap[i], heap[2*i+1] = heap[2*i+1], heap[i]
			i = 2*i + 1
		} else if 2*i+2 < len(heap) && x == heap[2*i+2] {
			heap[i], heap[2*i+2] = heap[2*i+2], heap[i]
			i = 2*i + 2
		} else {
			break
		}
	}

}

func max(a, b, c int) int {

	if a >= b && a >= c {
		return a
	} else if b >= a && b >= c {
		return b
	}
	return c
}
