package Assignment

func PairsWithGivenSum(A []int, B int) int {

	i := 0
	j := len(A) - 1

	count := 0
	for i < j {

		if A[i]+A[j] == B {
			count++
			i++
			j--
		} else if A[i]+A[j] < B {
			i++
		} else if A[i]+A[j] > B {
			j--
		}
	}

	return count
}
