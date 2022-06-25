package Assignment

func Lszero(A []int) []int {

	startIndex := -1
	endIndex := -1
	maxLength := 0

	countMap := make(map[int]int)
	prefixSumArray := make([]int, len(A))

	prefixSumArray[0] = A[0]

	for i := 1; i < len(A); i++ {
		prefixSumArray[i] = A[i] + prefixSumArray[i-1]
	}

	for i := range prefixSumArray {

		if prefixSumArray[i] == 0 {
			e := i
			s := 0
			length := e + 1
			if length > maxLength {
				maxLength = length
				startIndex = s
				endIndex = e
			}
			continue
		}

		if _, ok := countMap[prefixSumArray[i]]; ok {
			s := countMap[prefixSumArray[i]] + 1
			e := i
			length := e - s + 1
			if length > maxLength {
				maxLength = length
				startIndex = s
				endIndex = e
			}

		} else {
			countMap[prefixSumArray[i]] = i
		}
	}

	if startIndex == -1 {
		return []int{}
	}

	return A[startIndex : endIndex+1]
}
