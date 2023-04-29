package Assignment

import (
	"sort"
)

func NobleInteger(A []int) int {
	sort.Ints(A)

	ans := 0
	count := 0
	size := len(A) - 1

	if A[size] == 0 {
		ans++
		count++
	}

	for j := len(A) - 2; j >= 0; j-- {
		if A[j] != A[j+1] {
			count = size - j
		}
		if count == A[j] {
			ans++
		}
	}

	if ans == 0 {
		return -1
	}
	return ans

}
