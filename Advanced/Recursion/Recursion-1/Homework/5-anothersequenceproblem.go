package Homework

func AnotherSequenceProblem(A int) int {
	if A == 0 {
		return 1
	}
	if A > 0 && A < 3 {
		return A
	}

	return A + solve(A-1) + solve(A-2) + solve(A-3)

}
