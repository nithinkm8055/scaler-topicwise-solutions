package Assignment

func IsMagic(A int) int {
	return sump(A, 0)
}

// 1291
func sump(A, sum int) int {

	if A == 0 {
		if sum < 10 {
			if sum == 1 {
				return 1
			}
			return 0
		}

		A = sum
		sum = 0
	}

	return sump(A/10, sum+(A%10))
}
