package Assignment

func SubArrayZeroSum(A []int) int {

	countMap := make(map[int]int)
	prefixSumArray := make([]int, len(A))

	prefixSumArray[0] = A[0]

	for i := 1; i < len(A); i++ {
		prefixSumArray[i] = A[i] + prefixSumArray[i-1]
	}

	for i := range prefixSumArray {

		if prefixSumArray[i] == 0 {
			return 1
		}
		if _, ok := countMap[prefixSumArray[i]]; ok {
			return 1
		} else {
			countMap[prefixSumArray[i]] = 1
		}
	}

	return 0
}
