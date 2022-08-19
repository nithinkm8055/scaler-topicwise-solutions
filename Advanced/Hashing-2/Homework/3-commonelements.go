package Homework

func CommonElements(A []int, B []int) []int {

	countAMap := make(map[int]int)
	countBMap := make(map[int]int)

	for i := range A {
		countAMap[A[i]]++
	}

	for i := range B {
		countBMap[B[i]]++
	}

	result := make([]int, 0)

	for k, _ := range countAMap {
		if countAMap[k] > 0 && countBMap[k] > 0 {
			for i := 0; i < min(countAMap[k], countBMap[k]); i++ {
				result = append(result, k)
			}
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
