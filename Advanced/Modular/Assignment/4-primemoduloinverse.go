package Assignment

// fermats little theorem, a ^ (p-1) mod p = 1, where a < p and p is prime
//a inverse (mod p) is (a ^ (p-2)) mod p
func solve(A int, B int) int {
	return pow(A, B-2, B) % B
}

func power(A, B, C int) int {
	if A == 0 {
		return 0
	}

	if B == 0 {
		return 1
	}

	x := power(A, B/2, C)
	if B%2 == 0 {
		return (x * x) % C
	}

	return ((A % C) * ((x * x) % C)) % C
}
