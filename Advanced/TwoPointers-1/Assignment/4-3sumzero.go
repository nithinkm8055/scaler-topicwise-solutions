package Assignment

import "sort"

func threeSum(A []int) [][]int {

	sort.Ints(A)
	result := make([][]int, 0)
	for i := range A {

		s := -1 * A[i]

		j := i + 1
		k := len(A) - 1

		for j < k {
			if A[j]+A[k] == s {

				if len(result) > 0 {
					if !checkPresent(result, []int{A[i], A[j], A[k]}) {
						result = append(result, []int{A[i], A[j], A[k]})
					}
				} else {
					result = append(result, []int{A[i], A[j], A[k]})
				}

				j++
				k--

			} else if A[j]+A[k] < s {
				j++
			} else {
				k--
			}
		}

	}

	return result
}

func checkPresent(A [][]int, B []int) bool {
	for i := range A {
		count := 0
		for j := range A[i] {
			if A[i][j] != B[j] {
				break
			}
			count++
		}
		if count == 3 {
			return true
		}

	}
	return false
}
