package Homework

import (
	"sort"
)

func ElementsRemoval(A []int) int {
	sort.Ints(A)
	sum := 0
	for j := len(A) - 1; j >= 0; j-- {
		sum += A[j] * (len(A) - j)
	}
	return sum
}
