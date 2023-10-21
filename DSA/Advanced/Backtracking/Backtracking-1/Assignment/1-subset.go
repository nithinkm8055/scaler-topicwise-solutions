package Assignment

import "sort"

func subsets(A []int) [][]int {
	sort.Ints(A)
	temp := make([]int, 0)
	result := make([][]int, 0)
	subset(A, 0, len(A), temp, &result)
	return result
}

func subset(A []int, i int, n int, temp []int, result *[][]int) {

	if i == n {
		*result = append(*result, temp)
		return
	}

	subset(A, i+1, n, temp, result)
	temp = append(temp, A[i])
	subset(A, i+1, n, temp, result)

}
