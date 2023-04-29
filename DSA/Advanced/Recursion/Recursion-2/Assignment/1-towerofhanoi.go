package Assignment

func towerOfHanoi(A int) [][]int {

	result := make([][]int, 0)

	var towerOfHanoii func(int, int, int, int)

	towerOfHanoii = func(N, src, dest, temp int) {
		if N == 0 {
			return
		}

		towerOfHanoii(N-1, src, temp, dest)
		result = append(result, []int{N, src, dest})
		towerOfHanoii(N-1, temp, dest, src)

	}

	towerOfHanoii(A, 1, 3, 2)

	return result
}
