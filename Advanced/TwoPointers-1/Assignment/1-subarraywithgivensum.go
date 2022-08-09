package Assignment

func SubarrayWithGivenSum(A []int, B int) []int {

	pfSum := make([]int, len(A))

	pfSum[0] = A[0]

	if A[0] == B {
		return []int{A[0]}
	}

	for i := 1; i < len(A); i++ {
		pfSum[i] = pfSum[i-1] + A[i]

		if pfSum[i] == B {
			return A[:i+1]
		}
	}

	i := 1
	j := 1
	for j < len(A) {

		if pfSum[j]-pfSum[i-1] == B {
			return A[i : j+1]
		} else if pfSum[j]-pfSum[i-1] < B {
			j++
		} else if pfSum[j]-pfSum[i-1] > B {
			i++
		}
	}

	return []int{-1}
}
