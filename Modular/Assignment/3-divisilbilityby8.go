package Assignment

import (
	"fmt"
	"strconv"
)

func Divisibilityby8(A string) int {

	str := ""
	count := 1
	for j := len(A) - 1; j >= len(A)-3; j-- {
		if count <= len(A) {
			str = string(A[j]) + str
		}
		count++
	}
	atoi, _ := strconv.Atoi(str)
	fmt.Println(atoi)
	if atoi%8 == 0 {
		return 1
	}
	return 0
}
