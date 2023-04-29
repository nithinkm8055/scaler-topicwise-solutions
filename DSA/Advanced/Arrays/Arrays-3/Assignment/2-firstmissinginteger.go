package Assignment

import "fmt"

func FirstMissingPositive(A []int) int {
	i := 0
	for i < len(A) {
		if A[i] > 0 && A[i] <= len(A) {
			correctIndex := A[i] - 1
			if A[correctIndex] != A[i] {
				A[correctIndex], A[i] = A[i], A[correctIndex]
			} else {
				i++
			}
		} else {
			i++
		}
	}
	fmt.Println(A)
	for i := 0; i < len(A); i++ {
		if A[i] != i+1 {
			return i + 1
		}
	}

	return len(A)

}
