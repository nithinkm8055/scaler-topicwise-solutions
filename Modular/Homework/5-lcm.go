package Homework

func LCM(A int, B int) int {
	// lcm(a,b) * hcf(a,b) = a * b
	return (A * B) / HCF(A, B)
}

func HCF(a, b int) int {

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
