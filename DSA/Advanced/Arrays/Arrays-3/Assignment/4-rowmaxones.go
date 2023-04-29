package Assignment

func RowwithMaximumones(A [][]int) int {

	maxIndex := len(A[0]) - 1
	distance := -1
	ansIndex := -1

	for i := 0; i < len(A); i++ {

		for j := maxIndex; j >= 0; j-- {
			if A[i][j] == 1 {
				maxIndex = j

				currdistance := len(A[0]) - maxIndex
				if currdistance > distance {
					distance = currdistance
					ansIndex = i
				}

			}
		}
	}

	return ansIndex
}
