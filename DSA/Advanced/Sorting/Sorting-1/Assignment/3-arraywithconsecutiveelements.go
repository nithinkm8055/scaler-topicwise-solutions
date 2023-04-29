package Assignment

import "sort"

func ArrayWithConsecutivElements(A []int) int {

	sort.Ints(A)
	for i := 1; i < len(A); i++ {
		if A[i]-A[i-1] != 1 {
			return 0
		}
	}
	return 1
}
