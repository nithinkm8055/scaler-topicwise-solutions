package Assignment

func pow(A, B, C int) int {
	if A == 0 {
		return 0
	}

	if A < 0 {
		A = A + C
	}

	if B == 0 {
		return 1
	}

	x := pow(A, B/2, C)
	if B%2 == 0 {
		return (x * x) % C

	} else {

		return ((A % C) * ((x * x) % C)) % C
	}
}
