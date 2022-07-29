package Homework

import "sort"

func MaxMinMagic(A []int) []int {

	max := 0
	min := 0

	sort.Ints(A)

	for i := 0; i < len(A)/2; i++ {
		val := A[len(A)-1-i] - A[i]
		if val < 0 {
			val *= -1
		}
		max = ((max % (1000000007)) + (val % (1000000007))) % (1e9 + 7)
	}

	for j := len(A) - 1; j >= 0; j -= 2 {
		val := A[j] - A[j-1]
		min += val
	}

	return []int{max, min}
}
