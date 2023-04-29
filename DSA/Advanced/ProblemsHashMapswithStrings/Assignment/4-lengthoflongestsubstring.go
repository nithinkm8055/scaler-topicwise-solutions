package Assignment

func lengthOfLongestSubstring(A string) int {

	charMap := make(map[byte]int)

	l := 0
	r := 0
	ans := 0

	for r < len(A) {

		if _, ok := charMap[A[r]]; ok {
			ans = max(ans, len(charMap))
			for A[l] != A[r] {
				delete(charMap, A[l])
				l++
			}
			l++
			r++
		} else {
			charMap[A[r]]++
			r++
		}
	}
	ans = max(ans, len(charMap))

	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
