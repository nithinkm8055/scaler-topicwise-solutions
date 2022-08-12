package Assignment

func PairsWithGivenSumII(A []int, B int) int {

	i := 0
	j := len(A) - 1
	count := 0

	for i < j {

		if A[i]+A[j] == B {
			if A[i] == A[j] {
				freq := j - i + 1
				count += (freq * (freq - 1)) / 2
				i++
				j--
			} else {
				count1 := 1
				temp1 := A[i]
				i++

				for i < j && A[i] == temp1 {
					count1++
					i++
				}

				count2 := 1
				temp1 = A[j]
				j--
				for j >= 0 && A[j] == temp1 {
					j--
					count2++
				}

				count += (count1 * count2)
			}

		} else if A[i]+A[j] < B {
			i++
		} else {
			j--
		}
	}

	return count
}
