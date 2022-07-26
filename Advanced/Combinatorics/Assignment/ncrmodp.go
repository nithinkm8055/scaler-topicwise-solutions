package Assignment

func fact(n int, p int) int {
	if n == 0 {
		return 1
	}

	return ((n % p) * (fact(n-1, p) % p)) % p
}

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
	}
	return ((A % C) * ((x * x) % C)) % C
}

func ncrmodp(A int, B int, C int) int {

	a1 := fact(A, C) % C
	b1 := pow(fact(A-B, C), C-2, C) % C
	c1 := pow(fact(B, C), C-2, C) % C

	d1 := ((a1 % C) * (b1 % C)) % C
	ans := ((c1 % C) * (d1 % C)) % C

	return ans
}
