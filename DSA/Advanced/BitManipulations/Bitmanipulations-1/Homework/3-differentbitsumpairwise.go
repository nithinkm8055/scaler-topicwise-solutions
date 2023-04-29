package Homework

import "math"

func cntBits2(A []int) int {
	sum := 0

	for i := 0; i < 31; i++ {
		cnt := 0
		for j := 0; j < len(A); j++ {
			if A[j]&(1<<i) != 0 {
				cnt++
			}
		}
		sum += (cnt * (len(A) - cnt) * 2)
	}
	return sum % (int((math.Pow10(9))) + 7)
}

// TLE
func cntBits(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {

			xor := A[i] ^ A[j]
			cnt := 0
			for i := 0; i < 31; i++ {
				if xor&(1<<i) != 0 {
					cnt++
				}
			}

			sum += (2 * cnt)
		}
	}
	return sum
}
