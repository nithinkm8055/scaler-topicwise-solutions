package Assignment

import (
	"strconv"
)

func addBinary(A string, B string) string {

	if len(A) > len(B) {
		diff := len(A) - len(B)
		for i := 0; i < diff; i++ {
			B = "0" + B
		}
	} else if len(B) > len(A) {
		diff := len(B) - len(A)
		for i := 0; i < diff; i++ {
			A = "0" + A
		}
	}

	carry := 0
	result := ""
	for i := len(A) - 1; i >= 0; i-- {
		a, _ := strconv.Atoi(string(A[i]))
		b, _ := strconv.Atoi(string(B[i]))

		rem := (a + b + carry) % 2
		itoa := strconv.Itoa(rem)
		result = itoa + result
		carry = (a + b + carry) / 2
	}
	if carry > 0 {
		itoa := strconv.Itoa(carry)
		result = itoa + result
	}

	return result
}
