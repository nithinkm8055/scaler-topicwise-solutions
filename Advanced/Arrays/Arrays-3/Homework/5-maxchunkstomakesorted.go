package Homework

func MaxChunksToGetSorted(A []int) int {
	mx := -1
	chunks := 0
	for i := range A {

		if A[i] > mx {
			mx = A[i]
		}
		if i == mx {
			chunks++
		}
	}

	return chunks
}
