package Homework

func SecondLargest(A []int) int {
	largest := -1
	slarge := -1

	if len(A) < 2 {
		return -1
	}

	for i := range A {
		if A[i] > largest {
			slarge = largest
			largest = A[i]
		} else if A[i] > slarge {
			slarge = A[i]
		}
	}

	return slarge

}
