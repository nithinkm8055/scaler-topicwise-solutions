package Homework

import "sort"

// TODO: do with Binary Search
func solve(A []int, B int) int {
	sort.Ints(A)

	count := 0
	for i := range A {
		count++
		if count == B {
			return A[i]
		}
	}
	return -1
}
