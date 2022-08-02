package Assignment

/**
 * @input A : Integer
 * @input B : Integer
 * @input C : Integer
 *
 * @Output Integer
 */
func solve(A int, B int, C int) int {

	l := 1
	r := A * min(B, C)
	ans := 0

	for l <= r {

		mid := (l + r) / 2

		lcm := (B * C) / gcd(B, C)
		cnt := mid/B + mid/C - mid/lcm

		if cnt == A {
			ans = mid
			r = mid - 1
		} else if cnt > A {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return ans % (1000000007)
}

func gcd(a, b int) int {

	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a = a % b
		a, b = b, a
	}

	return a
}

func min(A, B int) int {
	if A < B {
		return A
	}
	return B
}
