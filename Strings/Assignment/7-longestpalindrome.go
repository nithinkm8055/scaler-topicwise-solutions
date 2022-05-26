package Assignment

func LongestPalindrome(A string) string {
	ans := ""
	max := 0
	temp, length := "", 0

	for i := range A {

		temp, length = isPalindrome(A, i, i)
		if length > max {
			ans = temp
			max = length
		}
		temp, length = isPalindrome(A, i, i+1)

		if length > max {
			ans = temp
			max = length
		}
	}

	return ans
}

func isPalindrome(A string, i int, j int) (string, int) {
	ans := ""
	length := 0
	for i >= 0 && j < len(A) {
		if A[i] == A[j] {
			ans = A[i : j+1]
			i--
			j++
		} else {
			break
		}
	}

	length = len(ans)

	return ans, length
}
