package Homework

func CpFact(A int, B int) int {

	// for 2 numbers to have a gcd of 1, either 1 of the number is 1
	// or one of the number is prime number, (the nos. cannnot be equal)

	// find largest prime divisor of A

	for i := 1; i <= A; i++ {
		if A%(A/i) == 0 {
			if gcd(B, A/i) == 1 {
				return A / i
			}
		}
	}
	return 0

}

func gcd(a int, b int) int {
	if a == 0 || b == 0 {
		if a > b {
			return a
		}
		return b
	}

	dividend := b
	divisor := a
	if a > b {
		dividend = a
		divisor = b
	}
	for divisor > 0 {
		remainder := dividend % divisor
		if remainder > 0 {
			dividend = divisor
			divisor = remainder
		} else {
			break
		}
	}
	return divisor
}
