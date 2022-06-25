package Assignment

func TitleToNumber(A string) int {
	sum := 0
	prod := 1
	for i := len(A) - 1; i >= 0; i-- {
		sum = sum + (int(A[i])-64)*prod
		prod = prod * 26
	}

	return sum
}
