package Assignment

func Size2Subsequence(A string) string {

	bytes := []byte(A)

	min := byte(127)
	minIndex := -1

	for i := 0; i < len(A)-1; i++ {
		if bytes[i] < min {
			minIndex = i
			min = bytes[i]
		}

	}

	secMin := byte(127)
	for j := minIndex + 1; j < len(A); j++ {
		if bytes[j] <= secMin {
			secMin = bytes[j]
		}
	}

	return string(min) + string(secMin)

}
