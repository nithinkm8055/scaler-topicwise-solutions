package Assignment

func findAthFibonacci(A int) int {

	if A < 2 {
		return A
	}

	return findAthFibonacci(A-1) + findAthFibonacci(A-2)
}
