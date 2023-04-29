package Homework

import (
	"strconv"
)

func Colorful(A int) int {
	// cannot be colorful if
	// contains 1
	// factors
	//
	itoa := strconv.Itoa(A)
	countMap := make(map[int]int)

	for i := range itoa {
		product := 1
		for j := i; j < len(itoa); j++ {
			atoi, _ := strconv.Atoi(string(itoa[j]))
			product = product * atoi

			if _, ok := countMap[product]; ok {
				return 0
			} else {
				countMap[product] = 1
			}
		}
	}
	return 1

}
