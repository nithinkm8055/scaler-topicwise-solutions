package Assignment

func TwoSum(A []int, B int) []int {

	indexMap := make(map[int]int)
	minIndex := 10000000
	for i := range A {
		diff := B - A[i]
		if _, ok := indexMap[A[i]]; ok {
			if i < minIndex {
				minIndex = i
			}
			return []int{indexMap[A[i]] + 1, i + 1}
		} else {
			if _, ok := indexMap[diff]; !ok {
				indexMap[diff] = i
			}
		}
	}

	return []int{}

}
