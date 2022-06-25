package Assignment

func solve(A int) int {
	if A == 0 {
		return 0
	}

	return A%10 + solve(A/10)
}
