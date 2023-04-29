package Assignment

func solve(A int) int {
	for A > 9 {
		A = ismagic(A)
	}
	if A == 1 {
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
