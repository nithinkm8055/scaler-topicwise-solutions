package Discussion

func MaximumPositivity(A []int) []int {

	endIndex := -1
	maxCount := 0
	count := 0
	for i := range A {
		if A[i] >= 0 {
			count++
			if count > maxCount {
				maxCount = count
				endIndex = i
			}
		} else {
			count = 0
		}
	}
	if endIndex == -1 {
		return A
	}
	startIndex := endIndex - maxCount + 1

	return A[startIndex : endIndex+1]

}
