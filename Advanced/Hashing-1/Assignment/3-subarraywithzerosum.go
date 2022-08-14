package Assignment

func SubarrayWithZeroSum(A []int) int {

	pfSum := make([]int, len(A))
	countMap := make(map[int]int)

	pfSum[0] = A[0]

	for i := 1; i < len(A); i++ {
		pfSum[i] = A[i] + pfSum[i-1]
		if pfSum[i] == 0 {
			return 1
		}

		if _, ok := countMap[pfSum[i]]; ok {
			return 1
		} else {
			countMap[pfSum[i]] = 1
		}
	}

	return 0

}
