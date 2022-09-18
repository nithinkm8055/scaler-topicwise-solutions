package Homework

func solve(A []int, k int) int {
	start := A[0]
	end := A[0]

	ans := -1

	for i := range A {
		if A[i] <= start {
			start = A[i]
		} else if A[i] > end {
			end = A[i]
		}
	}

	for start <= end {
		mid := (start + end) / 2
		numSmaller := 0

		for i := 0; i < len(A); i++ {
			if A[i] <= mid {
				numSmaller++
			}
		}

		if numSmaller >= k {
			end = mid - 1
			ans = mid
		} else {
			start = mid + 1
		}
	}

	return ans
}

//func solve(A []int, B int) int {
//	sort.Ints(A)
//
//	count := 0
//	for i := range A {
//		count++
//		if count == B {
//			return A[i]
//		}
//	}
//	return -1
//}
