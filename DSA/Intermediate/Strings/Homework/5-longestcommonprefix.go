package Homework

import (
	"sort"
)

func LongestCommonPrefix(A []string) string {

	sort.SliceStable(A, func(i, j int) bool {
		return len(A[i]) < len(A[j])
	})

	ans := A[0][:]
	for i := 1; i < len(A); i++ {

		for j := len(A[0]) - 1; j >= 0; j-- {
			if ans != A[i][:len(ans)] {
				ans = A[0][:len(ans)-1]
			} else {
				break
			}
		}
	}

	return ans
}
