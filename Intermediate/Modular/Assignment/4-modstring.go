package Assignment

import "strconv"

func FindMod(A string, B int) int {
	sum := 0
	power10 := 1
	for j := len(A) - 1; j >= 0; j-- {
		atoi, _ := strconv.Atoi(string(A[j]))
		sum = (sum + (atoi%B)*(power10%B)) % B
		power10 = (10 % B * power10 % B) % B
	}

	return sum
}
