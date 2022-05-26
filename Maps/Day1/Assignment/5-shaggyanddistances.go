package Assignment

func ShaggyandDistances(A []int) int {
	min := 10000000

	countMap := make(map[int]int)

	for i := range A {
		if _, ok := countMap[A[i]]; ok {
			distance := i - countMap[A[i]]
			if distance < min {
				min = distance
			}
		} else {
			countMap[A[i]] = i
		}
	}

	if min == 10000000 {
		return -1
	}

	return min
}
