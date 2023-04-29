package Assignment

func CommonElements(A []int, B []int) []int {

	countMap := make(map[int]int)
	result := make([]int, 0)

	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			countMap[A[i]] += 1
		} else {
			countMap[A[i]] = 1
		}
	}

	for i := range B {
		if _, ok := countMap[B[i]]; ok && countMap[B[i]] > 0 {
			result = append(result, B[i])
			countMap[B[i]] = countMap[B[i]] - 1
		}
	}

	return result
}
