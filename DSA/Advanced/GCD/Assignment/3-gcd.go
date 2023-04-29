package Assignment

// TC: O(log(max(a,b))
func gcd(a int, b int) int {
	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a = a % b
		a, b = b, a
	}

	return a
}
