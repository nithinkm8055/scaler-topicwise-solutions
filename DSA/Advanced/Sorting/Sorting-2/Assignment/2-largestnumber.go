package Assignment

import "sort"
import "strconv"

func largestNumber(A []int) string {

	var result string
	sort.SliceStable(A, func(i, j int) bool {

		num1 := strconv.Itoa(A[i]) + strconv.Itoa(A[j])
		num2 := strconv.Itoa(A[j]) + strconv.Itoa(A[i])

		return num1 > num2

	})

	flag := false
	for i := range A {
		if A[i] > 0 && !flag {
			flag = true
		}
		if flag {
			result = result + strconv.Itoa(A[i])
		}
	}

	if result == "" {
		return "0"
	}

	return result
}
