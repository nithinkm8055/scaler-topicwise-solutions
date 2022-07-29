package Homework

/**
 * @input A : Integer array
 *
 * @Output Integer array.
 */
import "sort"

func subUnsort(A []int) []int {
	temp := make([]int, 0)
	temp = append(temp, A...)

	sort.Ints(temp)

	minIndex := -1
	maxIndex := -1
	for i := range A {
		if A[i] != temp[i] {

			if minIndex == -1 {
				minIndex = i
			}

			if maxIndex == -1 || i > maxIndex {
				maxIndex = i
			}
		}
	}

	if minIndex == -1 {
		return []int{-1}
	}

	return []int{minIndex, maxIndex}
}
