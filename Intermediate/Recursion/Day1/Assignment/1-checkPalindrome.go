package Assignment

func checkPalindrome(A string, s, e int) bool {
	if s >= e {
		return true
	}

	if A[s] == A[e] {
		return checkPalindrome(A, s+1, e-1)
	}
	return false
}
