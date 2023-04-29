package Homework

import "sort"

type pair struct {
	div int
	quo int
}

func BthSmallestPrimeFraction(A []int, B int) []int {

	var result []pair
	// create all pairs
	for i := 0; i < len(A); i++ {
		for j := i + 1; j < len(A); j++ {
			result = append(result, pair{A[i], A[j]})
		}
	}

	// use comparators
	sort.SliceStable(result, func(i, j int) bool {
		return (float32(result[i].div) / float32(result[i].quo)) < (float32(result[j].div) / float32(result[j].quo))
	})

	return []int{result[B-1].div, result[B-1].quo}

}
