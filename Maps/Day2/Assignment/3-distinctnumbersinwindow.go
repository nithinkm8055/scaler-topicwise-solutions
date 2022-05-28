package Assignment

func DNums(A []int, B int) []int {
	windowMap := make(map[int]int)
	result := make([]int, 0)

	for i := 0; i < B; i++ {
		windowMap[A[i]] = windowMap[A[i]] + 1
	}

	result = append(result, len(windowMap))

	for i := 1; i <= len(A)-B; i++ {

		// adding element
		if _, ok := windowMap[A[B+i-1]]; ok {
			windowMap[A[B+i-1]] = windowMap[A[B+i-1]] + 1
		} else {
			windowMap[A[B+i-1]] = 1
		}

		//removing element
		if _, ok := windowMap[A[i-1]]; ok {
			windowMap[A[i-1]] = windowMap[A[i-1]] - 1
			if windowMap[A[i-1]] == 0 {
				delete(windowMap, A[i-1])
			}
		}

		result = append(result, len(windowMap))

	}

	return result
}
