package Homework

func MinSwaps(A []int, B int) int {

	windowSize := 0 // total no. of good elements in array

	for i := range A {
		if A[i] <= B {
			windowSize++
		}
	}

	if windowSize == len(A) {
		return 0
	}

	badValue := 0 // we have to minimize the bad value
	for i := 0; i < windowSize; i++ {
		if A[i] > B {
			badValue++
		}
	}

	minSwaps := badValue // stores the min swaps needed

	for i := 1; i <= len(A)-windowSize; i++ {

		if A[i-1] > B {
			badValue--
		}
		if A[windowSize+i-1] > B {
			badValue++
		}

		if badValue < minSwaps {
			minSwaps = badValue
		}

	}

	return minSwaps
}
