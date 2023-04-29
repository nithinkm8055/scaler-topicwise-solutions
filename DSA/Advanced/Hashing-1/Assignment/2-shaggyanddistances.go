package Assignment

func solve(A []int) int {
	ans := 100000000000

	indexMap := make(map[int]int)
	for i := range A {

		if _, ok := indexMap[A[i]]; ok {

			if i-indexMap[A[i]] < ans {
				ans = i - indexMap[A[i]]
			}
			indexMap[A[i]] = i
		} else {
			indexMap[A[i]] = i
		}

	}

	if ans == 100000000000 {
		return -1
	}

	return ans
}
