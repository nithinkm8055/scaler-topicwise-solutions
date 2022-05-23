package Homework

import (
	"sort"
	"strconv"
)

func LargestNumber(A []int) string {

	str := make([]string, 0)
	for i := range A {
		itoa := strconv.Itoa(A[i])
		str = append(str, itoa)
	}

	sort.SliceStable(str, func(i, j int) bool {
		ij := str[i] + str[j]
		ji := str[j] + str[i]

		return ij > ji
	})

	result := ""
	for i := range str {
		result = result + str[i]
	}

	if result[0] == '0' {
		return "0"
	}

	return result
}
