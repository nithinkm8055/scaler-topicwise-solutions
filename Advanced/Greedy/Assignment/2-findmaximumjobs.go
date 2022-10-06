package Assignment

import "sort"

type Pair struct {
	start int
	end   int
}

func solve(A []int, B []int) int {

	var pairs []Pair

	for i := range A {
		pairs = append(pairs, Pair{A[i], B[i]})
	}

	sort.SliceStable(pairs, func(i, j int) bool {
		return pairs[i].end < pairs[j].end
	})

	count := 0
	prev := pairs[0]
	for i := range pairs {

		if i == 0 {
			count++
		} else {

			if pairs[i].start >= prev.end {
				count++
				prev = pairs[i]
			}
		}
	}

	return count
}
