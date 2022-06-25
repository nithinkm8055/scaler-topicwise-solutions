package Homework

func AmazingSubarray(A string) int {
	count := 0
	for i := range A {
		if contains(string(A[i])) {
			count = count + (len(A) - i)
		}
	}
	return count
}

func contains(s string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}

	for i := range vowels {
		if vowels[i] == s {
			return true
		}
	}
	return false
}