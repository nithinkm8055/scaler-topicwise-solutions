package Homework

import "math"

func CountofDivisors(A []int) []int {

	result := make([]int, 0)
	for i := range A {
		result = append(result, factors(A[i]))
	}

	return result

}

//TLE
func factors(A int) int {
	cnt := 0
	for i := 1; i <= int(math.Sqrt(float64(A))); i++ {
		if A%i == 0 {
			if A/i == i {
				cnt += 1
			} else {
				cnt += 2
			}

		}
	}

	return cnt
}
