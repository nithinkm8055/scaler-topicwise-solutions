package Assignment

func GoodPair(A []int, B int) int {
	countMap := make(map[int]int)
	for i := range A {

		if _, ok := countMap[A[i]]; ok {
			return 1
		} else {
			val := B - A[i]
			countMap[val] = i
		}
	}

	return 0
}
