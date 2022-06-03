package Assignment

func pow(A int, B int, C int) int {

	if A < 0 {
		A = A + C
	}

	if A == 0 {
		return 0
	}
	if B == 0 {
		return 1
	} else if B%2 == 0 {
		y := pow(A, B/2, C)
		return (y * y) % C
	} else {
		return ((A % C) * pow(A, B-1, C)) % C
	}
}
