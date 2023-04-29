package Assignment

func lszero(A []int) []int {
	pfSum := make([]int, len(A))
	pfSum[0] = A[0]

	distMap := make(map[int]int) // stores first occurence of a key

	ans := -1
	lastIndex := -1

	if pfSum[0] == 0 {
		ans = 1
		lastIndex = 0
	}
	distMap[pfSum[0]] = 0
	for i := 1; i < len(A); i++ {
		pfSum[i] = A[i] + pfSum[i-1]

		if pfSum[i] == 0 {
			if i+1 > ans {
				ans = i + 1
				lastIndex = i
			}
		}
		if _, ok := distMap[pfSum[i]]; ok {
			if i-distMap[pfSum[i]] > ans {
				ans = i - distMap[pfSum[i]]
				lastIndex = i
			}
		} else {
			distMap[pfSum[i]] = i
		}
	}

	if lastIndex == -1 {
		return []int{}
	}

	if lastIndex-ans+1 < 0 {
		return A[:lastIndex+1]
	}

	return A[lastIndex-ans+1 : lastIndex+1]
}
