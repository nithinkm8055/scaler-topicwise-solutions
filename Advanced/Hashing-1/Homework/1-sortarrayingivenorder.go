package Homework

import "sort"

func solve(A []int, B []int) []int {
	sort.Ints(A)
	C := make([]int, 0)

	freqMap := make(map[int]int)
	for i := range A {
		freqMap[A[i]] = freqMap[A[i]] + 1
	}

	for i := range B {

		for freqMap[B[i]] > 0 {
			C = append(C, B[i])
			freqMap[B[i]] = freqMap[B[i]] - 1
		}

	}

	for i := range A {
		if freqMap[A[i]] > 0 {
			C = append(C, A[i])
		}
	}

	return C
}
