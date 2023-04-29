package Homework

func solve(A int, B int) int {

	if A == 1 || B == 1 {
		return 0
	}

	p := solve(A-1, (B+1)/2)

	if (B+1)%2 == 0 {
		return p
	}
	return 1 - p
}
