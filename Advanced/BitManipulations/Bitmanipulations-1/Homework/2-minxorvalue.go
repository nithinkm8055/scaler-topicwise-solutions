package Homework

import (
	"sort"
)

func findMinXor(A []int) int {
	min := 10000000000
	sort.Ints(A)
	for i := 0; i < len(A)-1; i++ {

		xor := A[i] ^ A[i+1]

		if xor < min {
			min = xor
		}

	}

	// TLE
	// for i := range A {
	//     for j := i+1; j < len(A); j++ {
	//         xor := A[i] ^ A[j]

	//         if xor < min {
	//             min = xor
	//         }
	//     }
	// }
	return min

}
