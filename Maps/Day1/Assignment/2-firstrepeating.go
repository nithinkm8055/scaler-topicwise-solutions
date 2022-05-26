package Assignment

func FirstRepeatingElement(A []int) int {
	countMap := make(map[int]int)

	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			countMap[A[i]] += 1
		} else {
			countMap[A[i]] = 1
		}
	}

	for i := range A {
		if countMap[A[i]] > 1 {
			return A[i]
		}
	}

	return -1
}
