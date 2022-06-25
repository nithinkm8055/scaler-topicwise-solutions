package Assignment

func PersonsOutsideTemple(A int, B [][]int) []int {

	result := make([]int, A)

	for i := range B {
		// start index of result array stores the amount received
		// end+1 index stores the negation of amount
		result[B[i][0]-1] += B[i][2]
		if B[i][1] < A {
			result[B[i][1]] += -1 * B[i][2]
		}
	}

	for i := 1; i < len(result); i++ {
		result[i] = result[i] + result[i-1]
	}

	return result
}
