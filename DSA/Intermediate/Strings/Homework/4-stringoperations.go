package Homework

func StringOperations(A string) string {
	bytes := []byte(A)

	for i := range bytes {
		if contains(bytes[i]) {
			bytes[i] = '#'
		} else if bytes[i] >= 65 && bytes[i] <= 90 {
			bytes[i] = 0
		}
	}

	ans := make([]byte, 0)
	for i := range bytes {
		if bytes[i] != 0 {
			ans = append(ans, bytes[i])
		}
	}

	return string(ans) + string(ans)
}

func contains(a byte) bool {
	vowels := []byte{'a', 'e', 'i', 'o', 'u'}

	for i := range vowels {
		if vowels[i] == a {
			return true
		}
	}
	return false
}
