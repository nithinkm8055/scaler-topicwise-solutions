package Homework

func ClosestPairFromSortedArrays(A []int, B []int, C int) []int {
	i := 0
	j := len(B) - 1

	iIndex := 1000000000
	jIndex := 1000000000
	dist := 1000000000
	for i < len(A) && j >= 0 {

		if A[i]+B[j] == C {
			return []int{A[i], B[j]}
		} else if A[i]+B[j] < C {
			diff := C - (A[i] + B[j])

			if diff < dist {
				dist = diff
				iIndex = i
				jIndex = j
			} else if diff == dist {
				if i <= iIndex && j < jIndex {
					iIndex = i
					jIndex = j
				}
			}
			i++
		} else {
			diff := (A[i] + B[j]) - C

			if diff < dist {
				dist = diff
				iIndex = i
				jIndex = j
			}

			j--
		}

	}

	return []int{A[iIndex], B[jIndex]}
}
