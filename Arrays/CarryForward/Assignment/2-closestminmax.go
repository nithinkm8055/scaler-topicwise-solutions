package Assignment

func ClosestMinMax(A []int) int {

	min := A[0]
	max := A[0]

	for i := range A {
		if A[i] > max {
			max = A[i]
		}
		if A[i] < min {
			min = A[i]
		}
	}

	minIndex := -1
	maxIndex := -1
	ans := -1

	for i := range A {
		if A[i] == min {
			minIndex = i
		}
		if A[i] == max {
			maxIndex = i
		}

		if minIndex != -1 && maxIndex != -1 {
			distance := maxIndex - minIndex
			if distance < 0 {
				distance = distance * -1
			}
			distance += 1

			if ans == -1 || distance < ans {
				ans = distance
			}
		}
	}
	return ans

}
