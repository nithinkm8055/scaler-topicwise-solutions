package Homework

func MaxSet(A []int) []int {

	sum := 0
	ans := 0

	sIndex := 0
	eIndex := 0
	ansSIndex := -1
	ansEIndex := -1
	flag := true

	maxDistance := 0
	distance := 0

	// find max sum subarray
	// if 2 arrays have same sum, return one with max length
	// if multiple arrays have sum and length, return one with minIndex

	for i := range A {
		if A[i] >= 0 {

			if sum == 0 && flag {
				sIndex = i
				eIndex = i
				flag = false
			}

			sum += A[i]
			if sum >= ans {
				eIndex = i

				if eIndex-sIndex+1 > distance {
					distance = eIndex - sIndex + 1

					if maxDistance == 0 || distance > maxDistance || sum > ans {
						ansSIndex = sIndex
						ansEIndex = eIndex
						maxDistance = distance
					}
				}

				ans = sum
			}
		} else {
			sum = 0
			flag = true
			distance = 0
		}
	}

	if ansSIndex == -1 {
		return []int{}
	}

	return A[ansSIndex : ansEIndex+1]

}
