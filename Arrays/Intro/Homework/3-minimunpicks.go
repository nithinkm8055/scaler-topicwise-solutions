package Homework

func MinPicks(A []int) int {

	maxEven := -1000000000
	minOdd := 1000000000

	for i := range A {
		if A[i]%2 == 0 {
			if A[i] > maxEven {
				maxEven = A[i]
			}
		} else {
			if A[i] < minOdd {
				minOdd = A[i]
			}
		}
	}

	return maxEven - minOdd

}
