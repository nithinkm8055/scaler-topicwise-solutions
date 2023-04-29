package Assignment

func dNums(A []int, B int) []int {
	dNums := make([]int, 0)
	numMap := make(map[int]int)

	for i := 0; i < B; i++ {
		numMap[A[i]] = numMap[A[i]] + 1
	}

	dNums = append(dNums, len(numMap))
	for i := 1; i <= len(A)-B; i++ {

		// remove element
		if numMap[A[i-1]] > 0 {
			numMap[A[i-1]] = numMap[A[i-1]] - 1
			if numMap[A[i-1]] == 0 {
				delete(numMap, A[i-1])
			}
		}

		// add element
		numMap[A[B+i-1]] = numMap[A[B+i-1]] + 1
		dNums = append(dNums, len(numMap))
	}
	return dNums
}
