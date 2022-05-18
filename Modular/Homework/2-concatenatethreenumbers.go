package Homework

import (
	"sort"
)

func Concatenate3Numbers(A int, B int, C int) int {
	ints := []int{A, B, C}
	sort.Ints(ints)

	return 10000*ints[0] + 100*ints[1] + ints[2]
}
