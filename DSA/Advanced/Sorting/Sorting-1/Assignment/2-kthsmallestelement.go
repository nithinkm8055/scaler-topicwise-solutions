package Assignment

import "sort"

func KthSmallest(A []int, B int) int {

	sort.Ints(A)

	return A[B-1]
}
