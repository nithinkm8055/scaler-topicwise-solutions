package Homework

func IsValidSudoku(A []string) int {
	if !checkRow(A) || !checkColumn(A) || !CheckBlock(A) {
		return 0
	}

	return 1
}

func checkRow(A []string) bool {

	for i := 0; i < 9; i++ {
		freqMap := make(map[byte]int)
		for j := 0; j < 9; j++ {
			if _, ok := freqMap[A[i][j]]; ok && A[i][j] != '.' {
				return false
			} else {
				freqMap[A[i][j]] = 1
			}
		}
	}
	return true
}

func checkColumn(A []string) bool {
	for j := 0; j < 9; j++ {
		freqMap := make(map[byte]int)
		for i := 0; i < 9; i++ {
			if _, ok := freqMap[A[i][j]]; ok && A[i][j] != '.' {
				return false
			} else {
				freqMap[A[i][j]] = 1
			}
		}
	}
	return true
}

func CheckBlock(A []string) bool {
	for k := 0; k < 9; k++ {
		x := k / 3
		y := k % 3
		freqMap := make(map[byte]int)
		for i := 3 * x; i < 3*x+3; i++ {
			for j := y * 3; j < y*3+3; j++ {
				if _, ok := freqMap[A[i][j]]; ok && A[i][j] != '.' {
					return false
				} else {
					freqMap[A[i][j]] = 1
				}
			}
		}

	}
	return true
}
