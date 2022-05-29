package Homework

import (
	"sort"
)

func IsDictionary(A []string, B string) int {
	C := make([]string, 0)
	C = append(C, A...)

	indexMap := make(map[byte]int)

	for i := range B {
		indexMap[B[i]] = i
	}

	sort.SliceStable(A, func(i, j int) bool {
		if indexMap[A[i][0]] < indexMap[A[j][0]] {
			return true
		} else if indexMap[A[i][0]] == indexMap[A[j][0]] {

			if len(A[i]) >= len(A[j]) {
				for k := 1; k < len(A[i]); k++ {
					if indexMap[A[i][k]] > indexMap[A[j][k]] {
						return false
					}
				}
			} else {
				//flag := false
				for k := 1; k < len(A[i]); k++ {
					if indexMap[A[i][k]] > indexMap[A[j][k]] {
						return false
					}
				}
				return true
			}
		}
		return false
	})

	for i := range A {
		if A[i] != C[i] {
			return 0
		}
	}

	return 1

}
