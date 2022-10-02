package Assignment

func solve(A []int) []int {

	result := make([]int, 0)
	result = append(result, []int{-1, -1}...)

	if len(A) == 1 {
		return []int{-1}
	} else if len(A) == 2 {
		return result
	}

	heapify(A[0:3], 0)
	result = append(result, A[0]*A[1]*A[2])
	for i := 3; i < len(A); i++ {

		if A[i] > A[0] {
			A[0], A[i] = A[i], A[0]
			heapify(A[0:3], 0)

		}

		result = append(result, A[0]*A[1]*A[2])
	}

	return result

}

func heapify(heap []int, i int) {

	for (2*i)+1 < len(heap) {
		x := min(heap[i], heap[2*i+1], -1)
		if 2*i+2 < len(heap) {
			x = min(heap[i], heap[2*i+1], heap[2*i+2])
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

func min(a, b, c int) int {

	if a <= b && a <= c {
		return a
	} else if b <= a && b <= c {
		return b
	}
	return c
}
