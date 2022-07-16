package Assignment

func solve(A int) int {
	A = ismagic(A)
	if A >= 10 {
		return solve(A)
	} else if A <= 10 && (A == 1) {
		return 1
	}
	return 0
}

func ismagic(A int) int {
	if A == 0 {
		return 0
	}

	return A%10 + ismagic(A/10)
}
