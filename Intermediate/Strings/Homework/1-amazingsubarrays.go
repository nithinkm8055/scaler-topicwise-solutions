package Homework

func AmazingSubArrays(A string) int {
	count := 0
	for i := range A {

		if containsVowel(A[i]) {
			count = count + (len(A) - 1) - i + 1
		}
	}

	return count % 10003

}

func containsVowel(s uint8) bool {

	vowels := []uint8{'A', 'E', 'I', 'O', 'U'}

	for i := range vowels {
		if vowels[i] == s || vowels[i]+32 == s {
			return true
		}
	}

	return false
}
