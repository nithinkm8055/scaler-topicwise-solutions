package Homework

import (
	"math"
	"sort"
)

func solve(A []int) int {

	sort.Ints(A)

	totalMax := 0
	totalMin := 0
	for i := range A {
		totalMax += (A[i] * power(2, i))
		totalMin += (A[i] * power(2, len(A)-i-1))

	}
	sum := totalMax - totalMin
	return sum % (int(math.Pow10(9)) + 7)
}

func power(A, n int) int {

	if n == 0 {
		return 1
	}

	x := power(A, n/2)

	if n%2 == 0 {
		return (x * x) % int((math.Pow10(9))+7)
	}

	return (A % int((math.Pow10(9))+7) * x * x) % int((math.Pow10(9))+7)
}
