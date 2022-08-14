package Assignment

func longestConsecutive(A []int) int {

	entryMap := make(map[int]int)

	for i := range A {
		entryMap[A[i]] = 0
	}

	ans := -1

	for k, _ := range entryMap {

		if _, ok := entryMap[k-1]; !ok {
			count := 0
			for i := k; i < 1e6+1; i++ {
				if _, ok := entryMap[i]; ok {
					count++
					ans = max(ans, count)
				} else {
					break
				}
			}

		}

	}
	return ans

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
