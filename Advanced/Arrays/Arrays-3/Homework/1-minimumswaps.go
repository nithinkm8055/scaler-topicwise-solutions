package Homework

func MinSwap(A []int, B int) int {

	windowSize := 0

	for i := range A {
		if A[i] <= B {
			windowSize++
		}
	}

	goodValue := 0
	badValue := 0

	for i := 0; i < windowSize; i++ {
		if A[i] <= B {
			goodValue++
		} else {
			badValue++
		}
	}

	ans := badValue

	j := windowSize
	for i := 1; i <= len(A)-windowSize; {

		if A[i-1] > B {
			badValue--
		}

		if A[j] > B {
			badValue++
		}

		i++
		j++
		ans = min(ans, badValue)
	}

	return ans
}
