package Assignment

func towerOfHanoi(A int) [][]int {

	result := make([][]int, 0)

	var towerOfHanoii func(int, int, int, int)

	towerOfHanoii = func(N, A, C, B int) {
		if N == 0 {
			return
		}

		towerOfHanoii(N-1, A, B, C)
		result = append(result, []int{N, A, C})
		towerOfHanoii(N-1, B, C, A)

	}

	towerOfHanoii(A, 1, 3, 2)

	return result

}
