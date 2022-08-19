package Homework

// TODO: Resolve TLE
func CompareSortedArrays(A []int, B [][]int) []int {

	result := make([]int, 0)

	for i := range B {
		countMap := make(map[int]int)

		for k := B[i][0]; k <= B[i][1]; k++ {
			countMap[A[k]]++
		}

		countMap2 := make(map[int]int)
		for l := B[i][2]; l <= B[i][3]; l++ {
			countMap2[A[l]]++
		}

		if len(countMap) != len(countMap2) {
			result = append(result, 0)
		} else {
			flag := false
			for k, _ := range countMap {
				if countMap[k] != countMap2[k] {
					result = append(result, 0)
					flag = true
					break
				}
			}
			if !flag {
				result = append(result, 1)
			}

		}
	}
	return result
}
