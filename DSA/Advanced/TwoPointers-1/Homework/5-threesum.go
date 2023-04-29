package Homework

import "sort"

func ThreeSumClosest(A []int, B int) int {

	sort.Ints(A)

	dist := 1000000000
	ans := 0

	for i := range A {

		diff := B - A[i]

		j := i + 1
		k := len(A) - 1

		for j < k {

			if A[j]+A[k] == diff {
				ans = B
				return ans
			} else if A[j]+A[k] < diff {
				if (diff - (A[j] + A[k])) < dist {
					dist = diff - (A[j] + A[k])
					ans = A[i] + A[j] + A[k]
				}

				j++
			} else {
				if (A[j]+A[k])-diff < dist {
					dist = (A[j] + A[k]) - diff
					ans = A[i] + A[j] + A[k]
				}
				k--
			}
		}
	}

	return ans
}
