package Homework

func MajorityElementNBy3(A []int) int {
	countMap := make(map[int]int)

	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			countMap[A[i]] = countMap[A[i]] + 1
		} else {
			countMap[A[i]] = 1
		}
	}
	max := countMap[A[0]]
	key := -1
	for k, v := range countMap {
		if v >= max {
			max = v
			key = k
		}
	}

	if max > len(A)/3 {
		return key
	}

	return -1
}
