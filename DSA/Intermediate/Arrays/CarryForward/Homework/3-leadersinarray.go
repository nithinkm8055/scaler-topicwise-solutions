package Homework

func LeadersInArray(A []int) []int {
	result := make([]int, 0)

	result = append(result, A[len(A)-1])
	max := A[len(A)-1]
	for i := len(A) - 2; i >= 0; i-- {
		if A[i] > max {
			max = A[i]
			result = append(result, A[i])
		}
	}
	return result
}
