package Homework

func sumofdigits(A int) int {

	if A == 0 {
		return 0
	}

	return A%10 + sumofdigits(A/10)
}
