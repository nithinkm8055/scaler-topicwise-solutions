package Homework

func InterestingArray(A []int) string {

	sum := 0

	for i := range A {
		sum += A[i]
	}

	if sum%2 == 0 {
		return "Yes"
	}

	return "No"

}
