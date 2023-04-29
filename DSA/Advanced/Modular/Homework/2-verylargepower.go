package Homework

func VeryLargePower(A int, B int) int {
	factorial := 1

	for i := 2; i <= B; i++ {
		factorial = (factorial * i) % (1e9 + 6)
	}

	return pow(A, factorial, 1e9+7)
}

func pow(A, B, C int) int {

	if B == 0 {
		return 1
	}

	x := pow(A, B/2, C)
	if B%2 == 0 {
		return (x * x) % C
	}

	return ((A % C) * ((x * x) % C)) % C
}
