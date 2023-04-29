package Assignment

import "sort"

func PairWithGivenDifference(A []int, B int) int {

	i := 0
	j := 1

	sort.Ints(A)
	count := 0
	for i < len(A) && j < len(A) {

		if A[j]-A[i] == B {
			count++

			for k := i + 1; k < len(A); k++ {
				if A[i] == A[k] {
					i++
				}
			}

			for m := j + 1; m < len(A); m++ {
				if A[j] == A[m] {
					j++
				}
			}

			i++
			j++
		} else if A[j]-A[i] < B {
			j++
		} else if A[j]-A[i] > B {
			i++
			if i == j {
				j++
			}
		}
	}

	return count
}
