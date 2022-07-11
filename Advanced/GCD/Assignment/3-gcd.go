package Assignment

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
