package Assignment

func DiffPossible(A []int, B int) int {

	foundMap := make(map[int]int)
	for i := range A {
		target := A[i] - B
		foundMap[target] = i
	}

	for i := range A {
		if _, ok := foundMap[A[i]]; ok && foundMap[A[i]] != i {
			return 1
		}
	}

	return 0

}
