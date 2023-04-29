package Assignment

// TC: O(26 *n)
func PermutationsofAinB(A string, B string) int {
	windowSize := len(A)

	aMap := make(map[byte]int)

	for i := range A {
		aMap[A[i]]++
	}

	bMap := make(map[byte]int)

	ans := 0
	for i := 0; i < windowSize; i++ {
		bMap[B[i]]++
	}

	// compare the maps
	if compare(aMap, bMap) {
		ans++
	}

	for i := 1; i <= len(B)-len(A); i++ {
		bMap[B[i-1]]--
		if bMap[B[i-1]] == 0 {
			delete(bMap, B[i-1])
		}

		bMap[B[len(A)+i-1]]++
		if compare(aMap, bMap) {
			ans++
		}
	}

	return ans

}

func compare(a, b map[byte]int) bool {

	if len(a) != len(b) {
		return false
	}

	for k, _ := range a {
		if a[k] != b[k] {
			return false
		}
	}

	return true

}
