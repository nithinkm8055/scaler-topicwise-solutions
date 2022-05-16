package Assignment

func subarraySum(A []int) int64 {

	sum := 0
	for i := range A {
		indexOccurence := (len(A) - i) * (i + 1)
		sum += indexOccurence * A[i]
	}
	return int64(sum)
}
