package Homework

func solve(A []int) []int {
	product := 1
	for i := range A {
		product *= A[i]
	}

	result := make([]int, len(A))
	for i := range A {
		result = append(result, product/A[i])
	}

	return result
}