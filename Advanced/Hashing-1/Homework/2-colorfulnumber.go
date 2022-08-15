package Homework

import "strconv"

func colorful(A int) int {
	itoa := strconv.Itoa(A)
	countMap := make(map[int]int)

	for i := range itoa {
		product := 1
		for j := i; j < len(itoa); j++ {
			atoi, _ := strconv.Atoi(string(itoa[j]))
			product = product * atoi
			countMap[product] = countMap[product] + 1
			if countMap[product] > 1 {
				return 0
			}
		}
	}
	return 1
}
